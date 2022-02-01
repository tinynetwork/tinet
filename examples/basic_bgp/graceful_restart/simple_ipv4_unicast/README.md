# Simple IPv4 unicast BGP-GR

```
$ tinet up -c spec.yaml | sudo sh
$ docker exec -it R2 bash
R2# tcpdump -nnli net0 # another shell
$ tinet conf -c spec.yaml | sudo sh 
$ docker exec R1 pkill -9 bgpd
```

