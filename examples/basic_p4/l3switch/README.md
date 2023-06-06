# P4Runtime and BM-CLI mapping

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
