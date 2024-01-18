#!/usr/bin/env python3
import os
import time
import struct
import pprint
import subprocess
import threading
from fcntl import ioctl
import p4runtime_sh.shell as sh
from p4runtime_sh.p4runtime import P4RuntimeClient
from p4.v1 import p4runtime_pb2
from p4.config.v1 import p4info_pb2


def exec(cmd):
    subprocess.check_call(cmd, shell=True)

def openTun(tunName, macaddr, ipaddr):
    tun = open("/dev/net/tun", "r+b", buffering=0)
    LINUX_IFF_TAP = 0x0002
    LINUX_IFF_NO_PI = 0x1000
    LINUX_TUNSETIFF = 0x400454CA
    flags = LINUX_IFF_TAP | LINUX_IFF_NO_PI
    ifs = struct.pack("16sH22s", tunName.encode("utf-8"), flags, b"")
    ioctl(tun, LINUX_TUNSETIFF, ifs)
    subprocess.check_call(f'ip link set {tunName} address {macaddr}', shell=True)
    subprocess.check_call(f'ip link set {tunName} up', shell=True)
    subprocess.check_call(f'ip addr add {ipaddr} dev {tunName}', shell=True)
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
te = sh.TableEntry("MyIngress.ipv4_lpm")(action="MyIngress.to_controller")
te.match["hdr.ipv4.dstAddr"] = "10.0.1.1/32"
te.insert()
te = sh.TableEntry("MyIngress.ipv4_lpm")(action="MyIngress.to_controller")
te.match["hdr.ipv4.dstAddr"] = "10.0.2.1/32"
te.insert()
te = sh.TableEntry("MyIngress.ipv4_lpm")(action="MyIngress.to_controller")
te.match["hdr.ipv4.dstAddr"] = "10.0.3.1/32"
te.insert()
sh.teardown()

## PREPARE TAP
swp = [
  openTun("swp1", "52:54:00:00:00:01", "10.0.1.1/24"),
  openTun("swp2", "52:54:00:00:00:02", "10.0.2.1/24"),
  openTun("swp3", "52:54:00:00:00:03", "10.0.3.1/24"),
]
exec('ip nei replace 10.0.1.2 lladdr 10:10:10:10:10:10 dev swp1')
exec('ip nei replace 10.0.2.2 lladdr 20:20:20:20:20:20 dev swp2')
exec('ip nei replace 10.0.3.2 lladdr 30:30:30:30:30:30 dev swp3')

## MAIN ROUTING
client = P4RuntimeClient(
  device_id=0,
  grpc_addr='localhost:9559',
  election_id=(0, 1))

def loop_packet_in():
    while True:
        rep = client.get_stream_packet("packet", timeout=1)
        if rep is not None:
            v = struct.unpack("@c", rep.packet.metadata[0].value)
            v = int.from_bytes(v[0], "little")
            print(f"PacketIN({v})")
            swp[v-1].write(rep.packet.payload)

def loop_packet_out(portIdx):
    port = swp[portIdx-1]
    while True:
        data = os.read(port.fileno(), 1500)
        print(f"PacketOut({portIdx})")
        req = p4runtime_pb2.StreamMessageRequest()
        req.packet.payload = data
        # XXX(slankdev)
        # metadata = p4runtime_pb2.PacketMetadata()
        # metadata.metadata_id = 1
        # metadata.value = portIdx
        # req.packet.metadata.append(metadata)
        # metadata.metadata_id = 3
        # metadata.value = mcast_grp
        # req.packet.metadata.append(metadata)
        client.stream_out_q.put(req)

thread1 = threading.Thread(target=loop_packet_in)
thread1.start()
threading.Thread(target=loop_packet_out, args=(1,)).start()
threading.Thread(target=loop_packet_out, args=(2,)).start()
threading.Thread(target=loop_packet_out, args=(3,)).start()
thread1.join()
