# the initial part is used to build cri-o and is based on
# https://github.com/kubernetes-incubator/cri-o/blob/v1.8.4/Dockerfile

FROM golang:1.8

ENV CRIO_TAG v1.8.4

# libseccomp in jessie is not _quite_ new enough -- need backports version
RUN echo 'deb http://httpredir.debian.org/debian jessie-backports main' > /etc/apt/sources.list.d/backports.list

RUN apt-get update && apt-get install -y \
    apparmor \
    autoconf \
    automake \
    bison \
    build-essential \
    curl \
    e2fslibs-dev \
    gawk \
    gettext \
    iptables \
    pkg-config \
    libaio-dev \
    libcap-dev \
    libfuse-dev \
    libostree-dev \
    libprotobuf-dev \
    libprotobuf-c0-dev \
    libseccomp2/jessie-backports \
    libseccomp-dev/jessie-backports \
    libtool \
    libudev-dev \
    protobuf-c-compiler \
    protobuf-compiler \
    python-minimal \
    libglib2.0-dev \
    libapparmor-dev \
    btrfs-tools \
    libdevmapper1.02.1 \
    libdevmapper-dev \
    libgpgme11-dev \
    liblzma-dev \
    netcat \
    socat \
    --no-install-recommends \
    && apt-get clean

# install bats
RUN cd /tmp \
    && git clone https://github.com/sstephenson/bats.git \
    && cd bats \
    && git reset --hard 03608115df2071fff4eaaff1605768c275e5f81f \
    && ./install.sh /usr/local

# install criu
ENV CRIU_VERSION 1.7
RUN mkdir -p /usr/src/criu \
    && curl -sSL https://github.com/xemul/criu/archive/v${CRIU_VERSION}.tar.gz | tar -v -C /usr/src/criu/ -xz --strip-components=1 \
    && cd /usr/src/criu \
    && make install-criu \
    && rm -rf /usr/src/criu

# Install runc
ENV RUNC_COMMIT 84a082bfef6f932de921437815355186db37aeb1
RUN set -x \
	&& export GOPATH="$(mktemp -d)" \
	&& git clone https://github.com/opencontainers/runc.git "$GOPATH/src/github.com/opencontainers/runc" \
	&& cd "$GOPATH/src/github.com/opencontainers/runc" \
	&& git fetch origin --tags \
	&& git checkout -q "$RUNC_COMMIT" \
	&& make static BUILDTAGS="seccomp selinux" \
	&& cp runc /usr/bin/runc \
	&& rm -rf "$GOPATH"

# Install CNI plugins
ENV CNI_COMMIT dcf7368eeab15e2affc6256f0bb1e84dd46a34de
RUN set -x \
       && export GOPATH="$(mktemp -d)" \
       && git clone https://github.com/containernetworking/plugins.git "$GOPATH/src/github.com/containernetworking/plugins" \
       && cd "$GOPATH/src/github.com/containernetworking/plugins" \
       && git checkout -q "$CNI_COMMIT" \
       && ./build.sh \
       && mkdir -p /opt/cni/bin \
       && cp bin/* /opt/cni/bin/ \
       && rm -rf "$GOPATH"

# Install custom CNI bridge test plugin
# XXX: this plugin is meant to be a replacement for the old "test_plugin_args.bash"
# we need this in testing because sandbox_run now gather IP address and the mock
# plugin wasn't able to properly setup the net ns.
# The bridge is based on the same commit as the one above.
#ENV CNI_COMMIT 6bfe036c38c8e1410f1acaa4b2ee16f1851472e4
ENV CNI_TEST_BRANCH custom-bridge
RUN set -x \
       && export GOPATH="$(mktemp -d)" \
       && git clone https://github.com/runcom/plugins.git "$GOPATH/src/github.com/containernetworking/plugins" \
       && cd "$GOPATH/src/github.com/containernetworking/plugins" \
       && git checkout -q "$CNI_TEST_BRANCH" \
       && ./build.sh \
       && mkdir -p /opt/cni/bin \
       && cp bin/bridge /opt/cni/bin/bridge-custom \
       && rm -rf "$GOPATH"

# Install crictl
ENV CRICTL_COMMIT b42fc3f364dd48f649d55926c34492beeb9b2e99
RUN set -x \
       && export GOPATH="$(mktemp -d)" \
       && git clone https://github.com/kubernetes-incubator/cri-tools.git "$GOPATH/src/github.com/kubernetes-incubator/cri-tools" \
       && cd "$GOPATH/src/github.com/kubernetes-incubator/cri-tools" \
       && git checkout -q "$CRICTL_COMMIT" \
       && go install github.com/kubernetes-incubator/cri-tools/cmd/crictl \
       && cp "$GOPATH"/bin/crictl /usr/bin/ \
       && rm -rf "$GOPATH"

# Make sure we have some policy for pulling images
RUN mkdir -p /etc/containers

RUN mkdir -p /go/src/github.com/kubernetes-incubator
RUN git clone -b "${CRIO_TAG}" https://github.com/kubernetes-incubator/cri-o /go/src/github.com/kubernetes-incubator/cri-o
RUN make -C /go/src/github.com/kubernetes-incubator/cri-o binaries crio.conf && \
    mkdir -p /tmp/cri-o/usr/bin && \
    cp /usr/bin/runc /tmp/cri-o/usr/bin && \
    DESTDIR=/tmp/cri-o make -C /go/src/github.com/kubernetes-incubator/cri-o install.bin install.config && \
    tar -C /tmp/cri-o -cvzf /crio.tar.gz .

FROM golang:1.8

# The following is based on https://github.com/kubernetes/release/blob/master/debian/Dockerfile

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -y && \
    apt-get -yy -q install --no-install-recommends --no-install-suggests --fix-missing \
      dpkg-dev build-essential debhelper dh-systemd socat && \
    apt-get upgrade -y && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    go get github.com/tcnksm/ghr && \
    go get github.com/Masterminds/glide

COPY --from=0 /crio.tar.gz /crio.tar.gz

ADD . /go/src/github.com/Mirantis/criproxy
WORKDIR /go/src/github.com/Mirantis/criproxy

CMD ["./build-package.sh"]
