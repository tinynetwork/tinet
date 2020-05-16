#!/usr/bin/env python

from __future__ import absolute_import
from __future__ import print_function

import grpc
from google.protobuf.any_pb2 import Any

import gobgp_pb2
import gobgp_pb2_grpc
import attribute_pb2

_TIMEOUT_SECONDS = 1000


def run():
    channel = grpc.insecure_channel('192.168.10.1:50051')
    stub = gobgp_pb2_grpc.GobgpApiStub(channel)

    nlri = Any()
    nlri.Pack(attribute_pb2.IPAddressPrefix(
        prefix_len=24,
        prefix="192.168.10.0",
    ))
    origin = Any()
    origin.Pack(attribute_pb2.OriginAttribute(
        origin=2,  # INCOMPLETE
    ))
    next_hop = Any()
    next_hop.Pack(attribute_pb2.NextHopAttribute(
        next_hop="0.0.0.0",
    ))
    attributes = [origin, next_hop]

    stub.AddPath(
        gobgp_pb2.AddPathRequest(
            table_type=gobgp_pb2.GLOBAL,
            path=gobgp_pb2.Path(
                nlri=nlri,
                pattrs=attributes,
                family=gobgp_pb2.Family(afi=gobgp_pb2.Family.AFI_IP, safi=gobgp_pb2.Family.SAFI_UNICAST),
            )
        ),
        _TIMEOUT_SECONDS,
    )

if __name__ == '__main__':
    run()
