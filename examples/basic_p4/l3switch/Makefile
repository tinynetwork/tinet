c:
	docker cp main.py P4:/
	docker exec -it P4 python3 /main.py

d:
	docker cp main.p4 P4:/
	docker exec P4 p4c --target bmv2 --arch v1model /main.p4 --p4runtime-files p4info.txt -o /
	docker exec P4 simple_switch_grpc /main.json \
		-i 1@vm1 -i 2@vm2 -i 3@vm3 \
		--nanolog ipc:///tmp/bm-0-log.ipc \
		--log-console -L debug \
		--notifications-addr ipc:///tmp/bmv2-0-notifications.ipc \
		-- --cpu-port 255
