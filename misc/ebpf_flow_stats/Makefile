DEV1 := ens4
DEV2 := veth05c9680
bpf-build:
	clang -target bpf -O3 -g -c filter.bpf.c
bpf-attach: bpf-build bpf-detach
	tc qdisc add dev $(DEV1) clsact
	tc qdisc add dev $(DEV2) clsact
	tc filter add dev $(DEV1) ingress pref 100 bpf obj filter.bpf.o section tc-ingress
	tc filter add dev $(DEV2) ingress pref 100 bpf obj filter.bpf.o section tc-ingress
bpf-detach:
	tc qdisc del dev $(DEV1) clsact || true
	tc qdisc del dev $(DEV2) clsact || true
bpf-hotswap:
	clang -target bpf -O3 -g -c filter.bpf.c
	tc filter change dev $(DEV1) ingress pref 100 chain 0 handle 0x1 bpf obj filter.bpf.o section tc-ingress
	tc filter change dev $(DEV2) ingress pref 100 chain 0 handle 0x1 bpf obj filter.bpf.o section tc-ingress
