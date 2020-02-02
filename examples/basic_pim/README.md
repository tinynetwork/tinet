
# PIM Multicast Test

multicast test
```
iperf -u -s -B 239.1.1.5 -i 1
iperf -u -c 239.1.1.5 -i <interval> -T <ttl> -t <time>
iperf -u -c 239.1.1.5 -i 1 -T 10 -t 5
```

## when add/del mroute

add
```
[ns0]ubuntu-bionic:~/git/netlinkd:) nlsniff -g all
monitoring group(RTMGRP) is 0xffffffff ...
RTM_NEWROUTE f=0x0000 s=0000000000 p=0000000000 :: fmly=128 dl=32 sl=32 tos=0 tab=253 pro=17 scope=0 type=5 f=0x0
  0x000f RTA_TABLE        :: 253
  0x0002 RTA_SRC          :: 0.0.0.0
  0x0001 RTA_DST          :: 239.1.1.5
  0x0003 RTA_IIF          :: 1
  0x0009 RTA_MULTIPATH    :: unknown-fmt(rta_len=20,data=08000001...)
  0x0011 RTA_MFC_STATS    :: unknown-fmt(rta_len=28,data=00000000...)
  0x0017 RTA_EXPIRES      :: unknown-fmt(rta_len=12,data=00000000...)
```

del
```
[ns0]ubuntu-bionic:~/git/netlinkd:) nlsniff -g all
monitoring group(RTMGRP) is 0xffffffff ...
RTM_NEWROUTE f=0x0000 s=0000000000 p=0000000000 :: fmly=128 dl=32 sl=32 tos=0 tab=253 pro=17 scope=0 type=5 f=0x0
  0x000f RTA_TABLE        :: 253
  0x0002 RTA_SRC          :: 0.0.0.0
  0x0001 RTA_DST          :: 239.1.1.5
  0x0003 RTA_IIF          :: 1
  0x0009 RTA_MULTIPATH    :: unknown-fmt(rta_len=12,data=08000001...)
  0x0011 RTA_MFC_STATS    :: unknown-fmt(rta_len=28,data=00000000...)
  0x0017 RTA_EXPIRES      :: unknown-fmt(rta_len=12,data=00000000...)

RTM_DELROUTE f=0x0000 s=0000000000 p=0000000000 :: fmly=128 dl=32 sl=32 tos=0 tab=253 pro=17 scope=0 type=5 f=0x0
  0x000f RTA_TABLE        :: 253
  0x0002 RTA_SRC          :: 0.0.0.0
  0x0001 RTA_DST          :: 239.1.1.5
  0x0003 RTA_IIF          :: 1
  0x0009 RTA_MULTIPATH    :: unknown-fmt(rta_len=12,data=08000001...)
  0x0011 RTA_MFC_STATS    :: unknown-fmt(rta_len=28,data=00000000...)
  0x0017 RTA_EXPIRES      :: unknown-fmt(rta_len=12,data=00000000...)
```
