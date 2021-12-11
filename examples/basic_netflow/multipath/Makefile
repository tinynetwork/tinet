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
	docker exec R1 pmacct -e -p /tmp/collect.pipe
	docker exec R2 pmacct -e -p /tmp/collect.pipe
	docker exec R3 pmacct -e -p /tmp/collect.pipe
	docker exec R4 pmacct -e -p /tmp/collect.pipe
test-iperf:
	#docker exec -it C2 iperf3 -c 10.1.0.10 -P5 -t 5
	docker exec -it C2 iperf3 -c 10.1.0.10 -n 5G
test-vegeta:
	docker exec -it C2 sh -c "cat /vegeta.conf | vegeta attack -keepalive=false -duration=5s | tee results.bin | vegeta report"
	#docker exec -it C2 sh -c "cat /vegeta.conf | vegeta attack -duration=5s | tee results.bin | vegeta report"
