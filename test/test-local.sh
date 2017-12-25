#!/bin/bash

# perform deployment test without CircleCI

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

export DOCKER_VERSION=17.03.0-ce
curl -sSL -o "/tmp/docker-${DOCKER_VERSION}.tgz" "https://get.docker.com/builds/Linux/x86_64/docker-${DOCKER_VERSION}.tgz"
tar -xz -C /tmp -f "/tmp/docker-${DOCKER_VERSION}.tgz"
mv /tmp/docker/* /usr/bin

# docker rm -f portforward || true
# ${SCRIPT_DIR}/portforward.sh start

(cd "${SCRIPT_DIR}/.." && ./build-package.sh)

# SKIP_SNAPSHOT=1 CRIPROXY_DEB=/go/src/github.com/Mirantis/criproxy-nodeps_0.9.3-2_amd64.deb test/deployment-test.sh
