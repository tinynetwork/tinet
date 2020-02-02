
help:
	@echo help

build:
	make -C /home/vagrant/git/frr.slankdev
	sudo make -C /home/vagrant/git/frr.slankdev install

install:
	docker exec R1 rm -rf /usr/lib/frr
	docker cp /usr/lib/frr R1:/usr/lib/frr
	docker cp /usr/bin/vtysh R1:/usr/bin/vtysh
	\
	docker exec R2 rm -rf /usr/lib/frr
	docker cp /usr/lib/frr R2:/usr/lib/frr
	docker cp /usr/bin/vtysh R2:/usr/bin/vtysh

config_srv6_only:
	docker cp frr.conf.srv6.R1 R1:/etc/frr/frr.conf
	docker cp frr.conf.srv6.R2 R2:/etc/frr/frr.conf

stop:
	docker exec R1 /usr/lib/frr/frrinit.sh stop
	docker exec R2 /usr/lib/frr/frrinit.sh stop

start:
	docker exec R1 /usr/lib/frr/frrinit.sh start
	docker exec R2 /usr/lib/frr/frrinit.sh start

restart:
	docker exec R1 /usr/lib/frr/frrinit.sh restart
	docker exec R2 /usr/lib/frr/frrinit.sh restart

capture:
	docker exec R1 pkill tcpdump | true
	docker exec R2 pkill tcpdump | true
	docker exec -d R1 tcpdump -ni net0 -w /tmp/r1.pcap
	docker exec -d R2 tcpdump -ni net0 -w /tmp/r2.pcap

nocapture:
	docker exec R1 pkill tcpdump | true
	docker exec R2 pkill tcpdump | true
	docker cp R1:/tmp/r1.pcap /vagrant/r1.pcap
	docker cp R2:/tmp/r2.pcap /vagrant/r2.pcap

rere:
	make -C . build
	make -C . install
	make -C . reload

taillog_R1:
	while :; do \
		docker exec -it R1 touch /tmp/frr.log; \
		docker exec -it R1 chown frr.frr /tmp/frr.log; \
		docker exec -it R1 tail -f /tmp/frr.log -n0; \
		sleep 1 ; done
taillog_R2:
	while :; do \
		docker exec -it R2 touch /tmp/frr.log; \
		docker exec -it R2 chown frr.frr /tmp/frr.log; \
		docker exec -it R2 tail -f /tmp/frr.log -n0; \
		sleep 1 ; done

tn_reconf:
	tn reconf | sudo sh
	make -C . install
	make -C . config_srv6_only

sh_route_vrfs:
	@echo -e '\n[[R1-glb]]'
	@docker exec R1 ip route list
	@docker exec R1 ip -6 route list
	@echo -e '\n[[R1-vrf1]]'
	@docker exec R1 ip route list vrf vrf1
	@echo -e '\n[[R1-vrf2]]'
	@docker exec R1 ip route list vrf vrf2
	@echo -e '\n[[R1-vrf3]]'
	@docker exec R1 ip route list vrf vrf3
	@echo -e '\n[[R2-glb]]'
	@docker exec R2 ip route list
	@docker exec R2 ip -6 route list
	@echo -e '\n[[R2-vrf1]]'
	@docker exec R2 ip route list vrf vrf1
	@echo -e '\n[[R2-vrf2]]'
	@docker exec R2 ip route list vrf vrf2
	@echo -e '\n[[R2-vrf3]]'
	@docker exec R2 ip route list vrf vrf3
	@echo

