FROM peterevans/vegeta as vegeta
FROM pmacct/pmacctd:v1.7.6
RUN apt update -y && apt install -y tcpdump netcat iperf3 watch file xxd psutils vim
RUN apt install -y nfdump jq
RUN apt install -y curl
COPY --from=vegeta /bin/vegeta /bin/vegeta
ENTRYPOINT bash
