
## IPsec Example

- libreswan

```bash
> docker exec R1 ipsec status | grep "Total IPsec connections" -A5
000 Total IPsec connections: loaded 1, active 1
000
000 State Information: DDoS cookies not required, Accepting new IKE connections
000 IKE SAs: total(2), half-open(0), open(0), authenticated(2), anonymous(0)
000 IPsec SAs: total(2), authenticated(2), anonymous(0)
000
```
