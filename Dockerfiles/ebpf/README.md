# Usage

```
docker build -t demo:latest .
docker run -it --rm --privileged demo:latest bash
./build_and_attach.sh
bpftool map dump name pkt_counter_egr
bpftool map dump name pkt_counter_ing
bpftool map event_pipe name pkt_counter_eve
# make some traffic on eth0
```
