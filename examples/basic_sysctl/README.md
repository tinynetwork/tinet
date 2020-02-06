
# sysctl example
```
tinet up | sh

tinet test all | sh
echo R1 sysctl state
docker exec R1 sysctl net.ipv4.ip_forward
docker exec R1 sysctl net.ipv4.ip_forward_use_pmtu
docker exec R1 sysctl net.ipv6.conf.all.forwarding
docker exec R1 sysctl net.ipv6.conf.all.disable_ipv6
echo R2 sysctl state
docker exec R2 sysctl net.ipv4.ip_forward
docker exec R2 sysctl net.ipv4.ip_forward_use_pmtu
docker exec R2 sysctl net.ipv6.conf.all.forwarding
docker exec R2 sysctl net.ipv6.conf.all.disable_ipv6
```
