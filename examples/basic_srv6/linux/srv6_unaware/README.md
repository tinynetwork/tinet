
# SRv6 Unaware Evaluation

![](./topo.jpeg)

send icmp traffic including `emacs` at R1
```
R1# ping 10.0.0.2 -p `echo "emacs" | xxd -p`
R1# ping 10.0.0.2 -p 656d6163732c
PATTERN: 0x656d6163732c
PING 10.0.0.2 (10.0.0.2) 56(84) bytes of data.
64 bytes from 10.0.0.2: icmp_seq=1 ttl=64 time=0.038 ms
64 bytes from 10.0.0.2: icmp_seq=2 ttl=64 time=0.075 ms
...
```

check traffic at R2
```
root@R2:/# tcpdump -ni net0 -X icmp
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on net0, link-type EN10MB (Ethernet), capture size 262144 bytes
14:04:45.602478 IP 10.0.0.2 > 10.0.0.1: ICMP echo reply, id 79, seq 3, length 64
        0x0000:  c688 a57b 07bd b6e5 b767 af7b 0800 4500  ...{.....g.{..E.
        0x0010:  0054 4a54 0000 4001 1c53 0a00 0002 0a00  .TJT..@..S......
        0x0020:  0001 0000 35be 004f 0003 fdc6 665c 0000  ....5..O....f\..
        0x0030:  0000 2931 0900 0000 0000 730a 656d 6163  ..)1......s.emac
        0x0040:  730a 656d 6163 730a 656d 6163 730a 656d  s.emacs.emacs.em
        0x0050:  6163 730a 656d 6163 730a 656d 6163 730a  acs.emacs.emacs.
```

Config filter @F1
```
// Apply filter
make -C filter
docker exec F1 ip link set net0 xdp obj /filter/net0
docker exec F1 ip link set net1 xdp obj /filter/net1

// Reset filter
docker exec F1 ip link set net0 xdp off
docker exec F1 ip link set net1 xdp off
```

