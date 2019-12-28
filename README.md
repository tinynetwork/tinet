# tinet-go

![](https://github.com/ak1ra24/tinet-go/workflows/golangci/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/ak1ra24/tinet-go)](https://goreportcard.com/report/github.com/ak1ra24/tinet-go)

Go implement of [slankdev/tinet](https://github.com/slankdev/tinet).

:warning: Operation check is still

## Setup
```
wget https://github.com/ak1ra24/tinet-go/releases/download/v0.0.0/tinet-go_linux_amd64.tar.gz -P /tmp
tar zxvf /tmp/tinet-go_linux_amd64.tar.gz -C /usr/local/bin
```

## Usage
```
tn up -c spec.yaml | sudo sh -x
tn conf -c spec.yaml | sudo sh -x
tn test -c spec.yaml | sudo sh -x
tn down -c spec.yaml | sudo sh -x
```

## Command Options

```
➜ tn                                                         
tinet is network simultulator created by docker

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