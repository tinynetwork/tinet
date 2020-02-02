
# CLOS Topology

Practice of designing DCN. Following are principle of Design.
- using modernaized technology (such as BGP-unnumbered.)

Version
- 0.0.0: basic CLOS-network ([yaml](./spec.v0.0.0.yaml))
- 0.0.1: using BGP-unnumbered ([yaml](./spec.v0.0.1.yaml))
- 0.0.2: using ECMP anycast ([yaml](./spec.v0.0.2.yaml))
- 0.0.3: add VM and ToR nodes ([yaml](./spec.v0.0.3.yaml))
- 0.0.4: support multi-tenancy ([yaml](./spec.v0.0.4.yaml)) (**currentry version**)
- 0.0.5: support SRv6 network slicing ([yaml](./spec.v0.0.5.yaml))

**version v0.0.3**
![](./topo.v0.0.3.png)

**version v0.0.2**
![](./topo.v0.0.2.png)

**version v0.0.1**
![](./topo.v0.0.1.png)

**version v0.0.0**
![](./topo.v0.0.0.png)

references
- LINE-SRv6-DCN ENOG55 http://enog.jp/wp-content/uploads/2018/12/05_20190222_ENOG55_LINE.pdf
- Large Scale DC Network Design https://www.slideshare.net/MasayukiKobayashi/dc-66865243
- Good TiNET examples by MIYA-kun https://github.com/mi2428/netben
- LINE DCN Overview by Kobayashi-san 2018.10 https://www.slideshare.net/linecorp/ss-116867631
- About designing the LINE-NW from scrach by Kobayashi-san 2019.01 https://www.janog.gr.jp/meeting/janog43/application/files/7915/4823/1858/janog43-line-kobayashi.pdf
- OpenStack Summit Vancouver 2018 Recap LINE-verda https://engineering.linecorp.com/ja/blog/openstack-summit-vancouver-2018-recap-2-2

operation performance
```
v0.0.3 sh -c 'tn up | sudo sh'
real    0m29.151s
user    0m6.945s
sys     0m4.472s

v0.0.3 sh -c 'tn conf | sudo sh' frr-all
real    1m25.104s
user    0m6.110s
sys     0m4.263s

v0.0.3 sh -c 'tn conf | sudo sh' frr-zebra,bgpd
real    1m13.867s
user    0m6.239s
sys     0m4.143s
```
