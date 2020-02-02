
# VPNv4 per-VRF using SRv6
- Hiroki Shirokura <slankdev@coe.ad.jp>
- 2018.12.31

![img](./topo.jpeg)

```
$ cd tinet/projects/basic_srv6/l3vpn
$ tn up | sudo sh # create network nodes.
$ tn conf | sudo sh # execute configuration to each nodes
$ docker ps # you can check some nodes exist.
$ tn test p2p | sudo sh # execute point-to-point link's ping.
$ tn test vpn | sudo sh # execute L3VPN ping.
```
