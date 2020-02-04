# tinet-go

![test](https://github.com/tinynetwork/tinet/workflows/test/badge.svg) ![release](https://github.com/tinynetwork/tinet/workflows/release/badge.svg) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/tinynetwork/tinet)](https://goreportcard.com/report/github.com/tinynetwork/tinet)

Go implement of [slankdev/tinet](https://github.com/slankdev/tinet).

## Release
[release](https://github.com/tinynetwork/tinet/releases)

## Requirements
- docker
- OpenvSwitch
- graphviz (if you want to use `tn img`)

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
