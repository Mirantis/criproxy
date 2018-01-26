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

kubectl="${HOME}/.kubeadm-dind-cluster/kubectl"
dind_script="dind-cluster-v1.8.sh"
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

step "Downloading kubeadm-dind-cluster script"
rm -f "${dind_script}"
wget "https://raw.githubusercontent.com/Mirantis/kubeadm-dind-cluster/master/fixed/${dind_script}"
chmod +x "${dind_script}"

step "Starting kubeadm-dind-cluster"

# Uncomment to have the cluster cleaned up.  Currently we're not doing
# it as it slows down local debugging.
# "./${dind_script}" clean

# Use single-worker cluster so as to have all the pods w/o tolerations
# scheduled on kube-node-1
NUM_NODES=1 "./${dind_script}" up

step "Propagating criproxy deb to the node"
if [[ ${CRIPROXY_DEB_URL:-} ]]; then
  docker exec kube-node-1 /bin/bash -c "curl -sSL '${CRIPROXY_DEB_URL}' >/criproxy.deb"
else
  docker cp "${CRIPROXY_DEB}" kube-node-1:/criproxy.deb
fi

step "Copying cri-o files to the node"
docker exec -i kube-node-1 tar -C / -xvz </crio.tar.gz
docker cp "${TEST_DIR}"/crio.service kube-node-1:/etc/systemd/system

step "Setting up criproxy and cri-o on the node"
docker exec -i kube-node-1 /bin/bash -s <<EOF
set -o errexit
set -o nounset
set -o pipefail
set -o errtrace

# FIXME: scopeo-containers package contains just a few files which we should
# probably just include with criproxy
export DEBIAN_FRONTEND=noninteractive
apt-get update -qq
apt-get install -qqy software-properties-common libgpgme11
add-apt-repository -y ppa:projectatomic/ppa
apt-get update -qq
apt-get install -qqy skopeo-containers

systemctl enable crio
systemctl start crio

dpkg -i /criproxy.deb
mkdir /etc/systemd/system/criproxy.service.d
echo -e '[Service]\nExecStart=\nExecStart=/usr/bin/criproxy.sh -v 3 -logtostderr -connect /var/run/dockershim.sock,cri.o:/var/run/crio.sock -listen /run/criproxy.sock' >/etc/systemd/system/criproxy.service.d/10-crio.conf
systemctl daemon-reload
systemctl restart criproxy
EOF

step "Starting and verifying busybox pod using CRI-O"
# Interrupting 'kubectl run' after grep finishes will cause it to exit
# with a non-zero status, thus '|| true'. It can be considered
# successful if it displays the message though
if ! ("${kubectl}" run bbtest-crio --attach \
        --overrides='{"metadata": {"annotations":{"kubernetes.io/target-runtime":"cri.o"}}}' \
        --image=cri.o/docker.io/busybox \
        --restart=Never -- \
        /bin/sh -c 'while true; do echo "this-is-crio-pod"; sleep 1; done' || true) |
        grep --line-buffered -m 1 this-is-crio-pod; then
  error "Failed to verify bbtest-crio pod"
fi

if ! "${kubectl}" logs bbtest-crio | grep -q this-is-crio-pod; then
  error "kubectl logs failed on bbtest-crio pod or didn't get this-is-crio-pod in its output"
fi

if ! docker exec kube-node-1 crioctl pod list | grep bbtest-crio; then
  error "Failed to find bbtest-crio pod among CRI-O pods"
fi

if docker exec kube-node-1 docker ps -a | grep bbtest-crio; then
  error "Error: found CRI-O pod's container among docker containers"
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
  error "Error: found docker pod in CRI-O pod list"
fi

if ! docker exec kube-node-1 docker ps -a | grep bbtest-docker; then
  error "Didn't find docker pod's container among docker containers"
fi

step "Verifying pod listing"
if ! "${kubectl}" get pods | grep bbtest-crio; then
  error "Failed to verify bbtest-crio pod"
fi
if ! "${kubectl}" get pods | grep bbtest-docker; then
  error "Failed to verify bbtest-docker pod"
fi

step "Deleting bbtest-crio pod"
"${kubectl}" delete pod bbtest-crio
wait-for "bbtest-crio pod to be gone" pod-is-gone bbtest-crio

step "Deleting bbtest-docker pod"
"${kubectl}" delete pod bbtest-docker
wait-for "bbtest-docker pod to be gone" pod-is-gone bbtest-docker

step "Making sure the cluster is still ok"
"${kubectl}" get pods --all-namespaces -o wide

exit "${status}"
