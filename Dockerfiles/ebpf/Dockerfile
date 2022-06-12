FROM ubuntu:rolling

ARG LIBBPF_VERSION="0.8.0"
ARG IPROUTE2_VERSION="5.18.0"
ARG BPFTOOL_VERSION="6.8.0"

RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y \
	vim curl git gcc make flex bison clang-12 libbsd-dev libbfd-dev libcap-dev \
  libelf-dev gcc-multilib pkg-config linux-tools-`uname -r`

RUN ln -s /usr/bin/clang-12 /usr/bin/clang

WORKDIR /opt

ADD https://github.com/libbpf/libbpf/archive/refs/tags/v${LIBBPF_VERSION}.tar.gz .
RUN tar xvf v${LIBBPF_VERSION}.tar.gz
RUN cd libbpf-${LIBBPF_VERSION}/src && make install BUILD_STATIC_ONLY=1 && make install_pkgconfig
RUN rm -rf libbpf-${LIBBPF_VERSION} v${LIBBPF_VERSION}.tar.gz

ADD https://git.kernel.org/pub/scm/network/iproute2/iproute2.git/snapshot/iproute2-${IPROUTE2_VERSION}.tar.gz .
RUN tar xvf iproute2-${IPROUTE2_VERSION}.tar.gz
RUN cd iproute2-${IPROUTE2_VERSION} && ./configure --libbpf_force=on --libbpf_dir=/ && make install
RUN rm -rf iproute2-${IPROUTE2_VERSION} iproute2-${IPROUTE2_VERSION}.tar.gz

ADD filter.c .
ADD build_and_attach.sh .
ADD detach.sh .
