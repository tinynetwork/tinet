# Multipath Configuration

diagram
```
+-----------+
|    R1     |
+-----------+
(net0) (net1)
  |      |
(net0) (net0)
+----+ +----+
| R2 | | R3 ||
+----+ +----+
(net1) (net1)
  |      |
(net0) (net1)
+-----------+
|    R4     |
+-----------+
```

other multipath config snippets
```
ip route add 1.1.1.1/32 \
    nexthop via 10.0.0.1 weight 1 \
    nexthop via 10.0.0.2 weight 2

ip route add :: table 10 \
  nexthop weight 10 encap seg6 mode encap segs A::,B:: via fe80::1 dev eth0 \
  nexthop weight 20 encap seg6 mode encap segs C::,D:: via fe80::2 dev eth0

ip route add :: vrf vrf0 \
  nexthop weight 10 encap seg6 mode encap segs A::,B:: via fe80::1 dev eth0 \
  nexthop weight 20 encap seg6 mode encap segs C::,D:: via fe80::2 dev eth0
```
