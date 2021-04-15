#!/bin/sh -xe
IMG=slankdev/iproute2:5.8
docker build -t $IMG .
