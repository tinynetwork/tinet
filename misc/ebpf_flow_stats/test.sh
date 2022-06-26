#!/bin/sh
set -xe
ID=$(sudo docker run -td nicolaka/netshoot)
sudo ./reload.py
sudo docker exec -it $ID apk update
sudo docker exec -it $ID curl ifconfig.io
sudo docker exec -it $ID curl slank.dev
