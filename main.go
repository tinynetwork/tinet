package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	name        = "tinet"
	description = "tinet: Tiny Network"
	version     = "0.0.0"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Version = version
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
