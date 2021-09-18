
# MP-BGP VPNv4 per-VRF w/ MPLS

![](./topo.png)

references: configure example of vpnv4 as small set.
https://gist.github.com/hkwi/5c116f05667a3abf43c7456fae32a529

setup
```
$ tn upconf | sudo sh
```

check vpn routes on R1
```
$ docker exec R1 vtysh -c 'show bgp ipv4 vpn'
BGP table version is 1, local router ID is 10.255.0.1, vrf id 0
Status codes:  s suppressed, d damped, h history, * valid, > best, = multipath,
               i internal, r RIB-failure, S Stale, R Removed
Nexthop codes: @NNN nexthop's vrf id, < announce-nh-self
Origin codes:  i - IGP, e - EGP, ? - incomplete

   Network          Next Hop            Metric LocPrf Weight Path
Route Distinguisher: 65001:1
*> 20.1.0.0/24      0.0.0.0@6<         0         32768 ?
    UN=0.0.0.0 EC{100:1} label=80 type=bgp, subtype=5
Route Distinguisher: 65001:2
*> 20.3.0.0/24      0.0.0.0@7<         0         32768 ?
    UN=0.0.0.0 EC{100:2} label=81 type=bgp, subtype=5
Route Distinguisher: 65001:3
*> 20.5.0.0/24      0.0.0.0@8<         0         32768 ?
    UN=0.0.0.0 EC{100:1} label=82 type=bgp, subtype=5
Route Distinguisher: 65002:1
*> 20.2.0.0/24      10.0.0.2         0             0 65002 ?
    UN=10.0.0.2 EC{100:1} label=80 type=bgp, subtype=0
Route Distinguisher: 65002:2
*> 20.4.0.0/24      10.0.0.2         0             0 65002 ?
    UN=10.0.0.2 EC{100:2} label=81 type=bgp, subtype=0
Route Distinguisher: 65002:3
*> 20.6.0.0/24      10.0.0.2         0             0 65002 ?
    UN=10.0.0.2 EC{100:2} label=82 type=bgp, subtype=0

Displayed  6 routes and 6 total paths
```

check vrf's route on VRF1 on R1 (VPNv4 rt100:1)
```
docker exec R1 vtysh -c 'show ip route vrf vrf1'
Codes: K - kernel route, C - connected, S - static, R - RIP,
       O - OSPF, I - IS-IS, B - BGP, E - EIGRP, N - NHRP,
       T - Table, v - VNC, V - VNC-Direct, A - Babel, D - SHARP,
       F - PBR,
       > - selected route, * - FIB route


VRF vrf1:
C>* 20.1.0.0/24 is directly connected, net1, 00:01:37
B>* 20.2.0.0/24 [200/0] via 10.0.0.2, net0(vrf Default-IP-Routing-Table), label 80, 00:01:30
B>* 20.5.0.0/24 [200/0] is directly connected, net3(vrf vrf3), 00:01:37
```

