#!/bin/sh
set -xe
ID=$(sudo docker run -td nicolaka/netshoot)
sudo ./reload.py
sudo docker exec -it $ID apk update
