#!/bin/sh
set -xe
docker rm -f $(docker ps -a -q)
