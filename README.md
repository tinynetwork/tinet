# tinet

![test](https://github.com/tinynetwork/tinet/workflows/test/badge.svg) ![release](https://github.com/tinynetwork/tinet/workflows/release/badge.svg) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/tinynetwork/tinet)](https://goreportcard.com/report/github.com/tinynetwork/tinet)

An instant virtual network on your laptop with
light-weight virtualization. Here we introduce the
Container Network Simulation tools. Users can generate,
from the YAML configuration file, the script to build
the L2 container network. Quickstart guide is provided
in QUICKSTART.md. It is tested on Ubuntu 16.04 LTS and
later.

## Requirements
- Docker
- OpenvSwitch (optional)
- graphviz (optional)

## Quick Install

There is only linux_amd64 pre-built binary
```
curl -Lo /usr/bin/tinet https://github.com/tinynetwork/tinet/releases/download/v0.0.3/tinet.linux_amd64
chmod +x /usr/bin/tinet
tinet --version
```

for ubuntu user
```
sudo apt update
sudo apt install -y linux-image-extra-virtual
sudo reboot
```

upgrading the kernel
```
$ sudo apt list "linux-image-5.15.*-generic"
linux-image-5.15.0-33-generic/focal-updates,focal-security 5.15.0-33.34~20.04.1 amd64
$ sudo apt install linux-image-5.15.0-33-generic linux-modules-5.15.0-33-generic linux-modules-extra-5.15.0-33-generic
$ sudo reboot
```
```
$ sudo grep 'menuentry ' $(sudo find /boot -name "grub.cfg") | cut -f 2 -d "'" | nl -v 0
     0  Ubuntu
     1  Ubuntu, with Linux 5.15.0-33-generic
     2  Ubuntu, with Linux 5.15.0-33-generic (recovery mode)
     3  Ubuntu, with Linux 5.4.0-113-generic
     4  Ubuntu, with Linux 5.4.0-113-generic (recovery mode)
$ sudo grub-set-default 3
$ sudo reboot
```

## Build
```
git clone https://github.com/tinynetwork/tinet tinet && cd $_
docker run --rm -i -t -v $PWD:/v -w /v golang:1.12 go build
mv tinet /usr/local/bin
```

## Usage

```
tinet up -c spec.yaml | sudo sh -x
tinet conf -c spec.yaml | sudo sh -x
tinet test -c spec.yaml | sudo sh -x
tinet down -c spec.yaml | sudo sh -x
docker run -it --rm --privileged --net=container:R1 nicolaka/netshoot bash
```

## Command Options

```
# tinet
NAME:
   tinet - tinet: Tiny Network

USAGE:
   tinet [global options] command [command options] [arguments...]

VERSION:
   0.0.1 (rev:)

AUTHOR:
   ak1ra24 <marug4580@gmail.com>

COMMANDS:
   check    check config
   conf     configure Node from tinet config file
   down     Down Node from tinet config file
   exec     Execute Command on Node from tinet config file.
   img      visualize network topology by graphviz from tinet config file
   init     Generate tinet config template file
   ps       docker and netns process
   pull     Pull Node docker image from tinet config file
   reconf   Stop, remove, create, start and config
   reup     Stop, remove, create, start
   test     Execute test commands from tinet config file.
   up       create Node from tinet config file
   upconf   Create, start and config
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Contribute

Simply fork and create a pull-request. We'll try to respond in a timely fashion.

## Links

- [Command Line Usage Example](docs/command-line-usage-example.md)
- [YAML Format](docs/specification_yml.md)
