#!/bin/sh
set -xe
sudo mkdir -p /sys/fs/bpf/xdp/globals
sudo mfpctl bpf xdp attach common --netns L1 --interface net0 --name l1 -v --define ETH_DST=1 -f
sudo mfpctl bpf xdp attach common --netns N1 --interface net0 --name n1 -v --define DEBUG_FUNCTION_CALL --define ETH_DST=2 -f
sudo mfpctl bpf map set-auto -f map.yaml
