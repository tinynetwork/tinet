
# nftables study (MASQ)

check nft is enabled (m is OK)
```
# cat /boot/config-`uname -r` | grep CONFIG_NF_TABLES
CONFIG_NF_TABLES=m
CONFIG_NF_TABLES_INET=m
CONFIG_NF_TABLES_NETDEV=m
CONFIG_NF_TABLES_IPV4=m
CONFIG_NF_TABLES_ARP=m
CONFIG_NF_TABLES_IPV6=m
CONFIG_NF_TABLES_BRIDGE=m
```

- Good reference
	-	https://knowledge.sakura.ad.jp/22636/
	- https://www.slideshare.net/s1061123/nftables-the-next-generation-firewall-in-linux
