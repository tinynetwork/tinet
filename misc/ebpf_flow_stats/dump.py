#!/usr/bin/env python3
import json
import pprint
import socket
import ipaddress
import subprocess

cmd = 'sudo bpftool map dump name flow_stats'
res = subprocess.check_output(cmd.split())
datas = json.loads(res)
# pprint.pprint(datas)

stats = {}
for data in datas:
    for element in data['elements']:
        daddr = element['key']['daddr']
        daddr = str(ipaddress.IPv4Address(socket.htonl(daddr)))
        key = "{}:{}".format(daddr, element['key']['dport'])
        cnt = 0
        for value in element['values']:
            cnt += value['value']['cnt']
        stats[key] = cnt

pprint.pprint(stats)
