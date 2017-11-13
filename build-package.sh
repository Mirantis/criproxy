#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail
set -o errtrace

if ! hash glide 2>/dev/null; then
  go get github.com/Masterminds/glide 1>&2
fi

if [[ ! -d vendor ]]; then
  glide install --strip-vendor 1>&2
fi

go build 1>&2
dpkg-buildpackage -us -uc -b 1>&2

(cd ../ && tar -c criproxy*deb)

