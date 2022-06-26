#!/usr/bin/env python3
import json
import pprint
import socket
import ipaddress
import subprocess


def execute(cmd, nojson=False):
  res = subprocess.check_output(cmd.split())
  if nojson:
    return
  return json.loads(res)


links = execute("ip -j link")
for link in links:
  if link["ifname"].startswith("ens4") or link["ifname"].startswith("veth"):
    qdisc_configured = False
    for qdisc in execute(f"tc -j qdisc list dev {link['ifname']} clsact"):
      if qdisc["kind"] == "clsact" and \
        qdisc["handle"] == "ffff:" and \
        qdisc["parent"] == "ffff:fff1":
        qdisc_configured = True
        break
    if qdisc_configured:
      execute(f"tc qdisc del dev {link['ifname']} clsact", nojson=True)
      print(f"{link['ifname']}/qdisc deleted")
