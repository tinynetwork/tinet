# TRex

```
tinet upconf | sudo sh -xe
docker exec -it T2 ./t-rex-64 --astf --astf-server-only -f tcp_open.py --cfg server.yaml
docker exec -it T1 ./t-rex-64 --astf --astf-client-mask 0x1 -f tcp_open.py --cfg client.yaml
```
