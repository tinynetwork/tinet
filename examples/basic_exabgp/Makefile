
exabgp:
	exabgp ./exabgp.conf

PCAP_PATH=/vagrant
getcapture:
	tcpdump -ni net0 -w $(PCAP_PATH)/in.pcap &
	exabgp ./exabgp.conf &
	sleep 2
	killall tcpdump
	killall python3

show:
	docker exec R1 vtysh -c 'sh bgp ipv4 uni'

