#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail
set -o errtrace

api_version=1.8
version_str=
if hash kubectl 2>/dev/null; then
    version_str="$(kubectl version --short --client)"
elif hash hyperkube 2>/dev/null; then
    version_str="$(hyperkube --version)"
elif [[ -x /k8s/hyperkube ]]; then
    # kubeadm-dind-cluster
    version_str="$(/k8s/hyperkube --version)"
fi

if [[ ${version_str} =~ v1\.7\. ]]; then
    api_version=1.7
fi

exec /usr/bin/criproxy -apiVersion "${api_version}" "$@"
