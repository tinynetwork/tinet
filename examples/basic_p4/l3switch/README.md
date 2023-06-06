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
te = table_entry["MyIngress.ipv4_lpm"](action="MyIngress.ipv4_forward")
te.match["hdr.ipv4.dstAddr"] = "192.168.20.20/32"
te.action["dstAddr"] = '20:20:20:20:20:20'
te.action["port"] = "2"
te.insert

te = table_entry["MyIngress.ipv4_lpm"](action="MyIngress.ipv4_forward")
te.match["hdr.ipv4.dstAddr"] = "192.168.10.10/32"
te.action["dstAddr"] = '10:10:10:10:10:10'
te.action["port"] = "1"
te.insert

te = table_entry["MyIngress.ipv4_lpm"](action="MyIngress.ipv4_forward");
te.match["hdr.ipv4.dstAddr"] = "192.168.30.30/32";
te.action["dstAddr"] = '30:30:30:30:30:30';
te.action["port"] = "3";
te.insert;

for te in table_entry["MyIngress.ipv4_lpm"].read(): print(te)
```
