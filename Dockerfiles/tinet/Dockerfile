FROM slankdev/iproute2:5.8 as iproute

FROM centos:centos7
RUN yum -y install git sudo tcpdump nmap-ncat \
    iptables-services traceroute
COPY --from=iproute /usr/sbin/ip /usr/sbin/ip
COPY --from=iproute /usr/sbin/ss /usr/sbin/ss
COPY --from=iproute /usr/sbin/tc /usr/sbin/tc
