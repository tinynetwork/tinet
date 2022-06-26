#!/bin/sh
set -xe
ID=$(sudo docker run -td nicolaka/netshoot)
sudo ./reload.py
