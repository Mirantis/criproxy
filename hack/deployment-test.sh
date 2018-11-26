#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail
set -o errtrace

if [ $(uname) = Darwin ]; then
  readlinkf(){ perl -MCwd -e 'print Cwd::abs_path shift' "$1";}
else
  readlinkf(){ readlink -f "$1"; }
fi
TEST_DIR="$(cd $(dirname "$(readlinkf "${BASH_SOURCE}")"); pwd)"

K8S_VERSION="${K8S_VERSION:-1.9}"
CLEAN_DIND="${CLEAN_DIND:-}"
kubectl="${HOME}/.kubeadm-dind-cluster/kubectl"
dind_script="dind-cluster-v${K8S_VERSION}.sh"
crictl="crictl -r unix:///var/run/containerd/containerd.sock"
status=0

if [[ ! ${CRIPROXY_DEB_URL:-} && ! ${CRIPROXY_DEB:-} ]]; then
  echo "Must specify either CRIPROXY_DEB_URL or CRIPROXY_DEB" >&2
  exit 1
fi

function msg {
  local color="${1}"
  local text="${2}"
  shift 2
  if [ -t 2 ] ; then
    echo -e "\x1B[97m* \x1B[${color}m${text}\x1B[39m $*" >&2
  else
    echo "* ${text} $*" >&2
  fi
}

function step {
  msg 92 "$@"
}

function error {
  msg 91 "ERROR $*"
  status=1
}

function wait-for {
  local title="$1"
  local action="$2"
  local what="$3"
  shift 3
  step "Waiting for:" "${title}"
  while ! "${action}" "${what}" "$@"; do
    echo -n "." >&2
    sleep 1
  done
  echo "[done]" >&2
}

function pod-is-gone {
  local name="$1"
  if "${kubectl}" get pods "${name}" >&/dev/null; then
    return 1
  fi
}

function start-kdc {
  step "Downloading kubeadm-dind-cluster script from ${dind_script}"
  rm -f "${dind_script}"
  wget "https://raw.githubusercontent.com/Mirantis/kubeadm-dind-cluster/master/fixed/${dind_script}"
  chmod +x "${dind_script}"

  step "Starting kubeadm-dind-cluster"

  if [[ ${CLEAN_DIND} ]]; then
    "./${dind_script}" clean
  fi

  # Use single-worker cluster so as to have all the pods w/o tolerations
  # scheduled on kube-node-1
  # APISERVER_PORT is set explicitly to avoid dynamic allocation
  # of the port by kdc
  APISERVER_PORT="${APISERVER_PORT:-8080}" NUM_NODES=1 "./${dind_script}" up
}

if [[ ! ${SKIP_START_KDC:-} ]]; then
  start-kdc
fi

step "Propagating criproxy deb to the node"
if [[ ${CRIPROXY_DEB_URL:-} ]]; then
  docker exec kube-node-1 /bin/bash -c "curl -sSL '${CRIPROXY_DEB_URL}' >/criproxy.deb"
else
  docker cp "${CRIPROXY_DEB}" kube-node-1:/criproxy.deb
fi

step "Setting up criproxy and dockershim+containerd on the node"
docker exec -i kube-node-1 /bin/bash -s <<EOF
set -o errexit
set -o nounset
set -o pipefail
set -o errtrace

# kill off k8s pods after stopping kubelet
systemctl stop kubelet
${crictl} ps -q   | xargs -r ${crictl} stop -t 0
${crictl} ps -qa  | xargs -r ${crictl} rm
${crictl} pods -q | xargs -r ${crictl} stopp
${crictl} pods -q | xargs -r ${crictl} rmp

export DEBIAN_FRONTEND=noninteractive
dpkg -i /criproxy.deb
sed -i 's@CRI_OTHER=.*@CRI_OTHER=containerd.io:/var/run/containerd/containerd.sock@' /etc/default/criproxy
cat <<CONF >>/etc/containerd/config.toml
[plugins]
  [plugins.cri]
    stream_server_address = ""
CONF
systemctl restart criproxy containerd
systemctl start kubelet
EOF

step "Starting and verifying busybox pod using containerd"
# Interrupting 'kubectl run' after grep finishes will cause it to exit
# with a non-zero status, thus '|| true'. It can be considered
# successful if it displays the message though
if ! ("${kubectl}" run bbtest-containerd --attach \
        --overrides='{"metadata": {"annotations":{"kubernetes.io/target-runtime":"containerd.io"}}}' \
        --image=containerd.io/docker.io/busybox \
        --restart=Never -- \
        /bin/sh -c 'while true; do echo "this-is-containerd-pod"; sleep 1; done' || true) |
        grep --line-buffered -m 1 this-is-containerd-pod; then
  error "Failed to verify bbtest-containerd pod"
fi

if ! "${kubectl}" logs bbtest-containerd | grep -q this-is-containerd-pod; then
  error "kubectl logs failed on bbtest-containerd pod or didn't get this-is-containerd-pod in its output"
fi

if ! docker exec kube-node-1 ${crictl} pods | grep bbtest-containerd; then
  error "Failed to find bbtest-containerd pod among containerd pods"
fi

if docker exec kube-node-1 docker ps -a | grep bbtest-containerd; then
  error "Error: found containerd pod's container among docker containers"
fi

step "Starting and verifying busybox pod using docker"
# Interrupting 'kubectl run' after grep finishes will cause it to exit
# with a non-zero status, thus '|| true'. It can be considered
# successful if it displays the message though
if ! ("${kubectl}" run bbtest-docker --attach \
        --image=docker.io/busybox \
        --restart=Never -- \
        /bin/sh -c 'while true; do echo "this-is-docker-pod"; sleep 1; done' || true) |
        grep --line-buffered -q this-is-docker-pod; then
  error "Failed to verify bbtest-docker pod"
fi

if ! "${kubectl}" logs bbtest-docker | grep -q this-is-docker-pod; then
  error "kubectl logs failed on bbtest-docker pod or didn't get this-is-docker-pod in its output"
fi

if docker exec kube-node-1 crioctl pod list | grep bbtest-docker; then
  error "Error: found docker pod in containerd pod list"
fi

if ! docker exec kube-node-1 docker ps -a | grep bbtest-docker; then
  error "Didn't find docker pod's container among docker containers"
fi

step "Verifying pod listing"
if ! "${kubectl}" get pods | grep bbtest-containerd; then
  error "Failed to verify bbtest-containerd pod"
fi
if ! "${kubectl}" get pods | grep bbtest-docker; then
  error "Failed to verify bbtest-docker pod"
fi

step "Deleting bbtest-containerd pod"
"${kubectl}" delete pod bbtest-containerd
wait-for "bbtest-containerd pod to be gone" pod-is-gone bbtest-containerd

step "Deleting bbtest-docker pod"
"${kubectl}" delete pod bbtest-docker
wait-for "bbtest-docker pod to be gone" pod-is-gone bbtest-docker

step "Making sure the cluster is still ok"
"${kubectl}" get pods --all-namespaces -o wide

exit "${status}"
