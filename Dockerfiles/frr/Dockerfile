FROM frr-ubuntu20:latest
USER root
RUN mkdir -p /etc/frr
COPY daemons /etc/frr/daemons
RUN apt install -y vim tcpdump
