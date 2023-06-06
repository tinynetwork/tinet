# P4Runtime and BM-CLI mapping

## Running p4runtime-shell

```
docker exec P4 simple_switch_grpc /main.json -i 1@vm1 -i 2@vm2 -i 3@vm3 --nanolog ipc:///tmp/bm-0-log.ipc --log-console -L debug --notifications-addr ipc:///tmp/bmv2-0-notifications.ipc -- --cpu-port 255
// copy p4info.txt and main.json
docker run -it --net container:P4 --name p4r -v /tmp/P4runtime-nanoswitch/:/tmp/ myproj/p4rt-sh-dev /bin/bash
source $VENV/bin/activate
/p4runtime-sh/p4runtime-sh --grpc-addr 0.0.0.0:9559 --device-id 0 --election-id 0,1 --config /p4info.txt,/main.json
```

## Misc
```
tables
actions
```

## Modifying Default Table Entry

bm cli of `simple_switch`
```
table_set_default MyIngress.ipv4_lpm MyIngress.drop
```

p4runtime cli of `simple_switch_grpc`
```
te = table_entry["MyIngress.ipv4_lpm"](action="MyIngress.drop", is_default=True)
te.modify
```

## Inserting Table Entry

bm cli of `simple_switch`
```
table_add MyIngress.ipv4_lpm MyIngress.ipv4_forward 192.168.10.10/32 => 10:10:10:10:10:10 1
table_add MyIngress.ipv4_lpm MyIngress.ipv4_forward 192.168.20.20/32 => 20:20:20:20:20:20 2
table_add MyIngress.ipv4_lpm MyIngress.ipv4_forward 192.168.30.30/32 => 30:30:30:30:30:30 3
```

p4runtime cli of `simple_switch_grpc`
```
te1 = table_entry["MyIngress.ipv4_lpm"](action="MyIngress.ipv4_forward")
te1.match["hdr.ipv4.dstAddr"] = "192.168.20.20/32"
te1.action["dstAddr"] = '20:20:20:20:20:20'
te1.action["port"] = "2"
te1.insert()

te2 = table_entry["MyIngress.ipv4_lpm"](action="MyIngress.ipv4_forward")
te2.match["hdr.ipv4.dstAddr"] = "192.168.10.10/32"
te2.action["dstAddr"] = '10:10:10:10:10:10'
te2.action["port"] = "1"
te2.insert()

te3 = table_entry["MyIngress.ipv4_lpm"](action="MyIngress.ipv4_forward")
te3.match["hdr.ipv4.dstAddr"] = "192.168.30.30/32"
te3.action["dstAddr"] = '30:30:30:30:30:30'
te3.action["port"] = "3"
te3.insert()

te4 = table_entry["MyIngress.ipv4_lpm"](action="MyIngress.ipv4_forward")
te4.match["hdr.ipv4.dstAddr"] = "192.168.40.40/32"
te4.action["dstAddr"] = 'aa:aa:aa:aa:aa:aa'
te4.action["port"] = "255"
te4.insert()

for te in table_entry["MyIngress.ipv4_lpm"].read(): print(te)
```
