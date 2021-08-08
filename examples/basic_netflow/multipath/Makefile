show:
	@echo -n "R1:net0  "
	@docker exec -it R1 sh -c "ip -j -s link show dev net0 | jq '.[0] | [.stats64.rx.packets, .stats64.tx.packets] | @tsv' -r"
	@echo -n "R2:net0  "
	@docker exec -it R1 sh -c "ip -j -s link show dev net0 | jq '.[0] | [.stats64.rx.packets, .stats64.tx.packets] | @tsv' -r"
	@echo -n "R3:net0  "
	@docker exec -it R3 sh -c "ip -j -s link show dev net0 | jq '.[0] | [.stats64.rx.packets, .stats64.tx.packets] | @tsv' -r"
	@echo -n "R4:net0  "
	@docker exec -it R4 sh -c "ip -j -s link show dev net0 | jq '.[0] | [.stats64.rx.packets, .stats64.tx.packets] | @tsv' -r"
