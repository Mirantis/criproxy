#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail
set -o errtrace

TARGET_PKGS=(v1_9 v1_10)
K8S_TAGS=(v1.9.3 v1.10.0)
SUBDIRS=(pkg/kubelet/apis/cri/v1alpha1/runtime pkg/kubelet/apis/cri/runtime/v1alpha2)
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
  subdir="${SUBDIRS[${i}]}"
  mkdir -p "${script_dir}/${dir}"
  for file in "${FILES[@]}"; do
    url="https://raw.githubusercontent.com/kubernetes/kubernetes/${tag}/${subdir}/${file}"
    subpath="${dir}/${file}"
    echo >&2 "Downloading ${url} -> ${subpath}"
    curl -sSL "${url}" >"${script_dir}/${subpath}"
     # | sed 's/^package v1alpha2/package runtime/g'
    if [[ ${file} = "api.pb.go" ]]; then
      sed -i "s/^func init *() *{ *$/func RegisterCRI() {/" "${script_dir}/${subpath}"
    fi
  done
done
