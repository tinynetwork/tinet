package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	name        = "tinet"
	description = "tinet: Tiny Network"
)

var (
	Version = "0.0.1"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Version = Version
	app.Usage = description
	app.Authors = []*cli.Author{
		{
			Name:  "ak1ra24",
			Email: "marug4580@gmail.com",
		},
	}
	app.Commands = commands

	return app
}
