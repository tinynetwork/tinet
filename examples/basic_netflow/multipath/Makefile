linkstats:
	@echo -n "R1:net0  "
	@docker exec -it R1 sh -c "ip -j -s link show dev net0 | jq '.[0] | [.stats64.rx.packets, .stats64.tx.packets] | @tsv' -r"
	@echo -n "R2:net0  "
	@docker exec -it R1 sh -c "ip -j -s link show dev net0 | jq '.[0] | [.stats64.rx.packets, .stats64.tx.packets] | @tsv' -r"
	@echo -n "R3:net0  "
	@docker exec -it R3 sh -c "ip -j -s link show dev net0 | jq '.[0] | [.stats64.rx.packets, .stats64.tx.packets] | @tsv' -r"
	@echo -n "R4:net0  "
	@docker exec -it R4 sh -c "ip -j -s link show dev net0 | jq '.[0] | [.stats64.rx.packets, .stats64.tx.packets] | @tsv' -r"
flowstats:
	docker exec C3 pmacct -s -p /tmp/collect.pipe
reset:
	docker exec C3 pmacct -e -p /tmp/collect.pipe
test:
	docker exec C1 iperf3 -c 10.2.0.10   -P10 -n 10G
