FROM golang:1.8

# based on https://github.com/kubernetes/release/blob/master/debian/Dockerfile

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -y && \
    apt-get -yy -q install --no-install-recommends --no-install-suggests --fix-missing \
      dpkg-dev build-essential debhelper dh-systemd && \
    apt-get upgrade -y && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ADD . /go/src/github.com/Mirantis/criproxy
WORKDIR /go/src/github.com/Mirantis/criproxy

# ENTRYPOINT ["./build-package.sh"]
