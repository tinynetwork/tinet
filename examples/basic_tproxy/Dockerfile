FROM debian:stretch-slim

RUN apt-get update
RUN apt-get install -y gcc iproute2 iptables

ADD tproxy.c /tmp
RUN gcc -o /usr/local/bin/tproxy /tmp/tproxy.c

CMD tproxy
