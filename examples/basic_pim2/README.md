
# PIM Multicast Test

![](topo.jpeg)

debug commands (vtysh)
```
show ip mroute
show ip pim neighbor
show ip pim join
show ip pim state
show ip pim rp-info
show ip pim interface
```

multicast test
```
iperf -u -s -B 239.1.1.5 -i 1
iperf -u -c 239.1.1.5 -i <interval> -T <ttl> -t <time>
iperf -u -c 239.1.1.5 -i 1 -T 10 -t 5
```

## ip mroute result

```
docker exec -it R1 vtysh -c 'show ip pim rp-info'
docker exec -it R2 vtysh -c 'show ip pim rp-info'
docker exec -it R3 vtysh -c 'show ip pim rp-info'

docker exec -it R1 vtysh -c 'show ip pim nei'
docker exec -it R2 vtysh -c 'show ip pim nei'
docker exec -it R3 vtysh -c 'show ip pim nei'
```
```
ip mroute
```
