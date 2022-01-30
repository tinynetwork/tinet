FROM golang:stretch

RUN apt-get update
RUN apt-get install -y iproute2 iptables

WORKDIR /opt

ADD go.mod /opt
ADD go.sum /opt
ADD http_interceptor.go /opt
RUN go build -o /usr/local/bin/http-interceptor

CMD http-interceptor
