#!/bin/bash -xe

clang -target bpf -O3 -g -c filter.c
tc qdisc del dev eth0 clsact || true
tc qdisc add dev eth0 clsact
tc filter add dev eth0 ingress bpf obj filter.o section tc-ingress
tc filter add dev eth0 egress bpf obj filter.o section tc-egress
