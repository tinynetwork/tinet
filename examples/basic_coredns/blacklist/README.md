
# CoreDNS example (very simple blacklist)

![](topo.png)

- Blacklisted
	-	emacs.org
	- sublimetext.com

```
docker exec R1 nslookup slank.dev
docker exec R1 nslookup test1.ichihara.org
docker exec R1 nslookup test2.ichihara.org

docker exec R1 nslookup www.vim.org
docker exec R1 nslookup emacs.org
docker exec R1 nslookup sublimetext.com
```
