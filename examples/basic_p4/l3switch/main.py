#!/usr/bin/env python3
import os
import time
import struct
import pprint
import subprocess
from fcntl import ioctl
import p4runtime_sh.shell as sh
from p4runtime_sh.p4runtime import P4RuntimeClient

def openTun(tunName):
    tun = open("/dev/net/tun", "r+b", buffering=0)
    LINUX_IFF_TAP = 0x0002
    LINUX_IFF_NO_PI = 0x1000
    LINUX_TUNSETIFF = 0x400454CA
    flags = LINUX_IFF_TAP | LINUX_IFF_NO_PI
    ifs = struct.pack("16sH22s", tunName.encode("utf-8"), flags, b"")
    ioctl(tun, LINUX_TUNSETIFF, ifs)
    subprocess.check_call(f'ip link set {tunName} up', shell=True)
    return tun

sh.setup(
  device_id=0,
  grpc_addr='localhost:9559',
  election_id=(0, 1),
  config=sh.FwdPipeConfig('/p4c/p4info.txt', '/main.json'))
te1 = sh.TableEntry("MyIngress.ipv4_lpm")(action="MyIngress.ipv4_forward")
te1.match["hdr.ipv4.dstAddr"] = "10.0.1.2/32"
te1.action["dstAddr"] = '10:10:10:10:10:10'
te1.action["port"] = "1"
te1.insert()
te2 = sh.TableEntry("MyIngress.ipv4_lpm")(action="MyIngress.ipv4_forward")
te2.match["hdr.ipv4.dstAddr"] = "10.0.2.2/32"
te2.action["dstAddr"] = '20:20:20:20:20:20'
te2.action["port"] = "2"
te2.insert()
te3 = sh.TableEntry("MyIngress.ipv4_lpm")(action="MyIngress.ipv4_forward")
te3.match["hdr.ipv4.dstAddr"] = "10.0.3.2/32"
te3.action["dstAddr"] = '30:30:30:30:30:30'
te3.action["port"] = "3"
te3.insert()
te4 = sh.TableEntry("MyIngress.ipv4_lpm")(action="MyIngress.ipv4_forward")
te4.match["hdr.ipv4.dstAddr"] = "10.0.0.0/8"
te4.action["dstAddr"] = 'ff:ff:ff:ff:ff:ff'
#te4.action["dstAddr"] = 'aa:aa:aa:aa:aa:aa'
te4.action["port"] = "255"
te4.insert()
sh.teardown()

## PREPARE TAP
tap = openTun("hoge")

## MAIN ROUTING
client = P4RuntimeClient(
  device_id=0,
  grpc_addr='localhost:9559',
  election_id=(0, 1))
while True:
    rep = client.get_stream_packet("packet", timeout=1)
    if rep is not None:
        #print("PacketIN")
        #pprint.pprint(rep.packet.payload)
        tap.write(rep.packet.payload)
