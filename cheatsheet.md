# Cheatsheet
## Linux Networking

```
# ip route add 10.0.0.0/24 encap seg6 mode encap segs a::,b::,c::,d:: via 2001:db8::1
# ip route add 10.0.0.0/24 encap seg6 mode l2encap segs a::,b::,c::,d:: via 2001:db8::1
# ip route add 10.0.0.0/24 encap seg6 mode inline segs a::,b::,c::,d:: via 2001:db8::1

# ip -6 route add fc00::1/128 encap seg6local \
    action End                    via 2001:db8::1
    action End.X   nh6 fc00::1:1  via 2001:db8::1
    action End.T   table 100      via 2001:db8::1
    action End.DX2 oif lxcbr0     via 2001:db8::1
    action End.DX4 nh4 10.0.3.254 via 2001:db8::1
    action End.DT4 vrftable 100   via 2001:db8::1
    action End.DX6 nh6 fc00::1:1  via 2001:db8::1
    action End.DT6 table 100      via 2001:db8::1
```

## VPP Networking
```
vpp> sr policy add bsid cafe::10 next fc00:2::10 fib-table 0
vpp> sr steer l3 192.168.1.0/24 via bsid cafe::10 fib-table 10
vpp> sr steer l2 host-net2 via bsid cafe::10

vpp> set sr encaps source addr fc00:1::
vpp> sr localsid address fc00:1::   behavior end
vpp> sr localsid address fc00:1::10 behavior end.dt4 10
vpp> sr localsid address fc00:1::10 behavior end.dx2 host-net2
```
