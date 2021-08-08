FROM nginx
RUN apt -y update -y && apt -y install iproute2
RUN apt update -y && apt install -y tcpdump netcat iperf3 watch file xxd psutils vim
ENTRYPOINT bash
