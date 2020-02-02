
# ExaBGP test

```
tn upconf | sudo sh
docker_mount_netns R2 ns0
ip netns exec ns0 bash
make       //execute exabgp
make getcapture //execute exabgp and pcap.
```


