#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail
set -o errtrace

TARGET_PKGS=(v1_7 v1_8 v1_9)
K8S_TAGS=(v1.7.12 v1.8.8 v1.9.3)
FILES=(api.pb.go api.proto constants.go)

if [ $(uname) = Darwin ]; then
  readlinkf(){ perl -MCwd -e 'print Cwd::abs_path shift' "$1";}
else
  readlinkf(){ readlink -f "$1"; }
fi
script_dir="$(cd $(dirname "$(readlinkf "${BASH_SOURCE}")"); pwd)"

for ((i = 0; i < ${#TARGET_PKGS[@]}; i++)); do
  dir="pkg/runtimeapi/${TARGET_PKGS[${i}]}"
  tag="${K8S_TAGS[${i}]}"
  mkdir -p "${script_dir}/${dir}"
  for file in "${FILES[@]}"; do
    url="https://raw.githubusercontent.com/kubernetes/kubernetes/${tag}/pkg/kubelet/apis/cri/v1alpha1/runtime/${file}"
    subpath="${dir}/${file}"
    echo >&2 "Downloading ${url} -> ${subpath}"
    curl -sSL "${url}" >"${script_dir}/${subpath}"
    if [[ ${file} = "api.pb.go" ]]; then
      sed -i "s/^func init *() *{ *$/func RegisterCRI() {/" "${script_dir}/${subpath}"
    fi
  done
done
