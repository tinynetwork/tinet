#!/usr/bin/env python3
import json
import pprint
import socket
import ipaddress
import subprocess
import sys


def execute(cmd, nojson=False):
  res = subprocess.check_output(cmd.split())
  if nojson:
    return
  return json.loads(res)


flow_stats_exist = False
for em in execute("sudo bpftool map -j"):
    if em["name"] == "flow_stats":
        flow_stats_exist = True
        break

if not flow_stats_exist:
    print("nothing")
    sys.exit(0)

stats = {}
for data in execute("sudo bpftool map dump name flow_stats"):
    for element in data['elements']:
        daddr = element['key']['daddr']
        daddr = str(ipaddress.IPv4Address(socket.htonl(daddr)))
        key = "{}:{}".format(daddr, element['key']['dport'])
        cnt = 0
        for value in element['values']:
            cnt += value['value']['cnt']
        stats[key] = cnt
pprint.pprint(stats)
