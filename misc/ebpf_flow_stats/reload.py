#!/usr/bin/env python3
import json
import pprint
import socket
import ipaddress
import subprocess


def execute_json(cmd):
  res = subprocess.check_output(cmd.split())
  datas = json.loads(res)
  return datas


def execute(cmd):
  subprocess.check_output(cmd.split())


links = execute_json("ip -j link")
for link in links:
  if link["ifname"].startswith("ens4") or link["ifname"].startswith("veth"):
    qdisc_configured = False
    for qdisc in execute_json("tc -j qdisc list dev ens4 clsact"):
      if qdisc["kind"] == "clsact" and \
        qdisc["handle"] == "ffff:" and \
        qdisc["parent"] == "ffff:fff1":
        qdisc_configured = True
        break
    if not qdisc_configured:
      execute(f"tc qdisc add dev {link['ifname']} clsact")
      print(f"{link['ifname']}/qdisc configured")
    else:
      print(f"{link['ifname']}/qdisc unchanged")

    filter_configured = False
    for filter in execute_json(f"tc -j filter list dev {link['ifname']}"+
        " ingress pref 100 chain 0 handle 0x1 protocol all"):
      if filter["kind"] == "bpf" and "options" in filter:
        if filter["options"]["bpf_name"] == "filter.bpf.o:[tc-ingress]":
          filter_configured = True
          break
    if not filter_configured:
      execute(f"tc filter add dev {link['ifname']} ingress "+
        "pref 100 bpf obj filter.bpf.o section tc-ingress")
      print(f"{link['ifname']}/filter configured")
    else:
      print(f"{link['ifname']}/filter unchanged")
