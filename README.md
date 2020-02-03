# tinet-go

![test](https://github.com/tinynetwork/tinet/workflows/test/badge.svg) ![release](https://github.com/tinynetwork/tinet/workflows/release/badge.svg) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/tinynetwork/tinet)](https://goreportcard.com/report/github.com/tinynetwork/tinet)

Go implement of [slankdev/tinet](https://github.com/slankdev/tinet).

## Requirements

- docker
- graphviz (if you want to use `tn img`)

## Install and Build

to install
```
wget https://github.com/tinynetwork/tinet/releases/download/v0.0.0/tinet-go_linux_amd64.tar.gz -P /tmp
tar zxvf /tmp/tinet-go_linux_amd64.tar.gz -C /usr/local/bin
```

to build
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
tinet is network emulator created by docker

Usage:
  tn [command]

Available Commands:
  build       Generate a Docker bundle from the spec file
  check       check config
  completion  Generates shell completion scripts
  conf        Execute config-cmd in a running container
  down        Stop and remove containers
  exec        Execute a command in a running container
  help        Help about any command
  img         Generate topology png file by graphviz
  init        Generate template spec file
  print       print tinet config file
  ps          List services
  pull        Pull Service images
  reconf      Stop, remove, create, start and config
  reup        Stop, remove, create and start
  test        Execute tests
  up          Create and start containers
  upconf      Create, start and config
  version     show the tinet version

Flags:
  -h, --help   help for tn

Use "tn [command] --help" for more information about a command.

```

## Contribute

Simply fork and create a pull-request. We'll try to respond in a timely fashion.

## Links

- [Command Line Usage Example](docs/command-line-usage-example.md)
- [YAML Format](docs/specification_yml.md)
