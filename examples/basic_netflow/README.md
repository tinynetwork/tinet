# NetFlow/IPFIX using pmacctd

## pmacct
ref: https://github.com/linsomniac/pmacct/blob/master/EXAMPLES

```
pmacctd -P print -r 1 -i net0 -c src_host,dst_host
pmacctd -P memory -i net0 -c src_host,dst_host
pmacctd -P memory -c src_host,dst_host
pmacctd -P memory -c src_host,dst_host -D
pmacctd -P memory -c src_host,dst_host,proto,src_port,dst_port
pmacctd -f /conf.txt

pmacct -s -p /tmp/collect.pipe
pmacct -s -p /tmp/collect.pipe -O json
```

```
plugins:            memory
aggregate:          src_host,dst_host,proto,src_port,dst_port
plugin_buffer_size: 35200
plugin_pipe_size:   409600000
```

## nfcapd/nfdump
```
nfcapd -w -l /tmp -p 2100
```
