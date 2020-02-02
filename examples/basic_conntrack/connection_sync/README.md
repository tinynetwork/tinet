
# Connection Sync with conntrackd and keepalived

```
tn upconf | sudo sh
```

<!-- docker exec -it S1 tcpdump -ni net0 -Qin '(tcp[tcpflags] & tcp-syn)' != 0 -->
<!-- docker exec C1 curl --interface 10.0.0.2 20.0.0.2 -->
<!-- docker exec C1 curl --interface 10.0.0.3 20.0.0.2 -->
<!-- docker exec C1 curl --interface 10.0.0.4 20.0.0.2 -->
<!-- docker exec S1 conntrack -L -->
