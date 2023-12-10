# TRex

```
tinet upconf | sudo sh -xe
docker exec -it T1
./t-rex-64 -i --astf --cfg ./cfg.yaml
./trex-console
trex>start -f http_eflow.py -t cps=10
```
