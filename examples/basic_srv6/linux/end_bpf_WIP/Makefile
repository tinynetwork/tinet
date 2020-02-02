
all:
	clang -O2 -Wall -target bpf -c filter.c
	docker cp filter.o R2:/filter.o
	docker exec R2 \
		ip -6 route add fc00:2::20 encap seg6local action End.BPF \
		endpoint object /filter.o section prog dev lo

del:
	docker exec R2 \
		ip -6 route del fc00:2::20

