
# HV route advertisement study for BGP-DCN

- [GOOD reference about routing-info manupulation by @ukinau](https://engineering.linecorp.com/ja/blog/openstack-summit-vancouver-2018-recap-2-2/)
- [GOOD reference about proxy arp by @ukinau](https://qiita.com/ukinau/items/cb25588fb0c276a009dc)

![](topo.png)

```
docker exec TOR vtysh -c 'show bgp ipv4 unicast'
BGP table version is 4, local router ID is 10.255.0.99, vrf id 0
Status codes:  s suppressed, d damped, h history, * valid, > best, = multipath,
               i internal, r RIB-failure, S Stale, R Removed
Nexthop codes: @NNN nexthop's vrf id, < announce-nh-self
Origin codes:  i - IGP, e - EGP, ? - incomplete

   Network          Next Hop            Metric LocPrf Weight Path
*> 1.1.1.1/32       0.0.0.0                  0         32768 ?
*> 10.0.0.11/32     dn1                      0             0 65001 ?
*> 10.0.0.12/32     dn1                      0             0 65001 ?
*> 10.0.0.13/32     dn1                      0             0 65001 ?

Displayed  4 routes and 4 total paths
```

