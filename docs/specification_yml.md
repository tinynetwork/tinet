# Yaml Format

## Node Definition

### Node type

- name: node name. It will be container-name or netns-name.
- type: node type (default: docker)
	- docker: node is docker container
	- netns: node is just network namespace
- image: specify docker-image
- sysctls: set sysctls
- mounts: mounts file/directory on the container
- dns: set DNS resolver
- dns_search: set DNS search domain


```
nodes:
  - name: Node0
    image: ubuntu:18.04
  - name: Node1
    type: netns
	- name: Node2
	  build: .
```

### Interface Definition

- type
	- direct: p2p connect to other container
	- bridge: bridge connection
	- phys  : host's network interface
- addr: specify mac address

```
nodes:
  - name: Node0
    image: ubuntu:18.04
    interfaces:
      - { name: net0, type: direct, args: R0#net1 }
      - { name: net0, type: bridge, args: B0 }
      - { name: eth0, type: phys }
      - { name: net0, type: direct, args: R1#net1, addr: 11:11:11:11:11:11 }
```

### Bridge Definition

If you use the bridge interface type, you need to
define the Bridge-Instance. It'll be created as a
linux bridge instance.

- name: interface name
- type: interface type, you must choose following.
	- docker: net-if of docker container
	- netns: net-if of network namespace
	- phys: host's network interface

```
bridges:
  - name: Bridge0
    interfaces:
      - { name: net0, type: docker, args: R0 }
      - { name: net0, type: netns, args: NS0 }
      - { name: eth0, type: phys }
```

## Config Definition
- name: node name
- cmds: shell command

```
node_configs:
  - name: Router1
    cmds:
      - cmd: ip addr add 10.255.0.10/32 dev lo
      - cmd: ip addr add 10.0.0.10/24 dev net0
			- copy: ./config.conf /usr/local/config.conf
  - name: Router0
    cmds:
      - cmd: sysctl -w 'net.ipv4.fib_multipath_hash_policy=1'
      - cmd: /usr/lib/frr/frr start
      - cmd: ip addr add 10.255.0.1/32 dev lo
      - cmd: ip addr add 10.0.0.1/24 dev net0
      - cmd: ip addr add 10.1.0.1/24 dev net1
      - cmd: >-
          vtysh -c 'conf t'
          -c 'router bgp 100'
          -c ' bgp router-id 10.255.0.1'
          -c ' neighbor 10.0.0.10 remote-as 100'
          -c ' neighbor 10.0.0.20 remote-as 100'
          -c ' neighbor 10.0.0.30 remote-as 100'
```

## Test Definition

- name: test name, you can specify when you run `tn conf -n <name>`
- cmds: same format of config definition

```
test:
  - name: p2p
	  cmds:
    - cmd: docker exec S0 ping -c2 10.1.0.1
    - cmd: docker exec S1 ping -c2 192.168.0.1
    - cmd: docker exec S2 ping -c2 192.168.0.1
    - cmd: docker exec S3 ping -c2 192.168.0.1
  - name: remote
	  cmds:
    - cmd: docker exec S0 ping -c2 10.1.0.1
    - cmd: docker exec S1 ping -c2 192.168.0.1
    - cmd: docker exec S2 ping -c2 192.168.0.1
    - cmd: docker exec S3 ping -c2 192.168.0.1
```

## System Requirement Definition

```
require:
	- kernel_min: 4.11.0
	- kernel_max: 4.15.0
	- kmod: [ mpls_router, mpls_iptunnel, mpls_gso ]
  - kconfig: [ CONFIG_NET_L3_MASTER_DEV ]
```
