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
cd "$(dirname "$(readlinkf "${BASH_SOURCE}")")"/..

# FIXME: without --skip-unsafe, it may generate direct translations
# for FilesystemUsage in some cases
conversion-gen --logtostderr \
               -i github.com/Mirantis/criproxy/pkg/runtimeapis/v1_9 \
               --base-peer-dirs github.com/Mirantis/criproxy/pkg/runtimeapis/v1_9 \
               -h hack/boilerplate.go.txt \
               --skip-unsafe \
               github.com/Mirantis/criproxy/pkg/runtimeapis/v1_9

sed -i 's/^package v1_9/package runtime/' pkg/runtimeapis/v1_9/conversion_generated.go
