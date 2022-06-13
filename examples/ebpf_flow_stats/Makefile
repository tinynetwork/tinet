DEV := ens4
bpf-attach:
	clang -target bpf -O3 -g -c filter.bpf.c
	tc qdisc del dev $(DEV) clsact || true
	tc qdisc add dev $(DEV) clsact
	tc filter add dev $(DEV) ingress bpf obj filter.bpf.o section tc-ingress
	tc filter add dev $(DEV) egress bpf obj filter.bpf.o section tc-egress
bpf-detach:
	tc qdisc del dev $(DEV) clsact || true
