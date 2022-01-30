FROM debian:stretch-slim

RUN apt-get update
RUN apt-get install -y gcc iproute2 iptables

ADD http_interceptor.c /tmp
RUN gcc -o /usr/local/bin/http-interceptor /tmp/http_interceptor.c

CMD http-interceptor
