
NAME=N1
log_N1:
	while :; do \
		docker exec $(NAME) vtysh -c 'conf te' -c 'log file /tmp/frr.log'; \
		docker exec -it $(NAME) tail -f /tmp/frr.log; \
		sleep 1 ; done

