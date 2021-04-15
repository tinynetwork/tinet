# BUILDER_STAGE
FROM centos:centos7 as iproute
RUN yum -y remove iproute \
 && yum -y groupinstall "Development tools"
RUN yum -y install elfutils-devel elfutils-libelf elfutils-libelf-devel
ADD https://mirrors.edge.kernel.org/pub/linux/utils/net/iproute2/iproute2-5.8.0.tar.gz /tmp
RUN cd /tmp && tar xpf iproute2-5.8.0.tar.gz && cd iproute2-5.8.0 \
 && ./configure && make && make install

# FINAL_STAGE
FROM centos:centos7
RUN yum -y install \
	sudo iproute traceroute iputils bash-completion tcpdump \
	wget vim iptables-services less
COPY --from=iproute /usr/sbin/ip /usr/sbin/ip
COPY --from=iproute /usr/sbin/ss /usr/sbin/ss
COPY --from=iproute /usr/sbin/tc /usr/sbin/tc
