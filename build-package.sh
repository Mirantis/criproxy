#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail
set -o errtrace

if [[ ! -d vendor ]]; then
  glide install --strip-vendor 1>&2
fi

go build 1>&2

# https://www.debian.org/doc/manuals/maint-guide/update.en.html#idm3360
date="$(LANG=C date -R)"
version="$(git describe 2>/dev/null | sed 's/^v\|-g.*//g' || true)"
version="${version:-0.0.0}"
author="Ivan Shvedunov <ishvedunov@mirantis.com>"

cat >debian/changelog <<EOF
criproxy (${version}) stable; urgency=optional

  * https://github.com/Mirantis/criproxy

 -- ${author}  ${date}

EOF
dpkg-buildpackage -us -uc -b 1>&2

# (cd ../ && tar -c criproxy*deb)
