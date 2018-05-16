#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail
set -o errtrace

TARGET_PKGS=(v1_9 v1_10)
K8S_TAGS=(v1.9.7 v1.10.2)
SUBDIRS=(pkg/kubelet/apis/cri/v1alpha1/runtime pkg/kubelet/apis/cri/runtime/v1alpha2)
FILES=(api.pb.go api.proto constants.go)

if [ $(uname) = Darwin ]; then
  readlinkf(){ perl -MCwd -e 'print Cwd::abs_path shift' "$1";}
else
  readlinkf(){ readlink -f "$1"; }
fi
top_dir="$(cd "$(dirname "$(readlinkf "${BASH_SOURCE}")")"/..; pwd)"

for ((i = 0; i < ${#TARGET_PKGS[@]}; i++)); do
  dir="pkg/runtimeapis/${TARGET_PKGS[${i}]}"
  tag="${K8S_TAGS[${i}]}"
  subdir="${SUBDIRS[${i}]}"
  mkdir -p "${top_dir}/${dir}"
  for file in "${FILES[@]}"; do
    url="https://raw.githubusercontent.com/kubernetes/kubernetes/${tag}/${subdir}/${file}"
    subpath="${dir}/${file}"
    echo >&2 "Downloading ${url} -> ${subpath}"
    curl -sSL "${url}" >"${top_dir}/${subpath}"
    # | sed 's/^package v1alpha2/package runtime/g'
  done
done

cd "${top_dir}"
sed -i 's@^\(type StorageIdentifier struct\)@/* +k8s:conversion-gen=false */ \1@' pkg/runtimeapis/v1_9/api.pb.go
go fmt pkg/runtimeapis/v1_9/api.pb.go

tar --exclude='vendor' --exclude='.git' -c . |
    docker run --rm -i ishvedunov/criproxy-build:0.0.4 \
           /bin/bash -c 'cd /go/src/github.com/Mirantis/criproxy && tar -x && hack/generate.sh && tar -c $(find . -name "*_generated.go")' | tar -xv
