# TRex

```
tinet upconf | sudo sh -xe
./do.sh
docker exec -it -w /opt/trex/automation/trex_control_plane/interactive/trex/examples/astf T1 python3 new_connection_test.py
```

```
docker run -it --net container:HV1 nicolaka/netshoot tcpdump -qtnni any
```
