package main

import "github.com/urfave/cli/v2"

var commands = []*cli.Command{
	commandBuild,
	commandCheck,
	commandConf,
	commandDown,
	commandExec,
	commandImg,
	commandInit,
	commandPs,
	commandPull,
	commandReConf,
	commandReUp,
	commandTest,
	commandUp,
	commandUpConf,
}

var commandBuild = &cli.Command{
	Name:   "build",
	Usage:  "Build docker Image from tinet config file",
	Action: CmdBuild,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Verbose",
		},
	},
}

var commandUp = &cli.Command{
	Name:   "up",
	Usage:  "create Node from tinet config file",
	Action: CmdUp,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Verbose",
		},
	},
}

var commandConf = &cli.Command{
	Name:   "conf",
	Usage:  "configure Node from tinet config file",
	Action: CmdConf,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Verbose",
		},
	},
}

var commandUpConf = &cli.Command{
	Name:   "upconf",
	Usage:  "Create, start and config",
	Action: CmdUpConf,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Verbose",
		},
	},
}

var commandDown = &cli.Command{
	Name:   "down",
	Usage:  "Down Node from tinet config file",
	Action: CmdDown,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Verbose",
		},
	},
}

var commandExec = &cli.Command{
	Name:   "exec",
	Usage:  "Execute Command on Node from tinet config file.",
	Action: CmdExec,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
	},
}

var commandImg = &cli.Command{
	Name:   "img",
	Usage:  "visualize network topology by graphviz from tinet config file",
	Action: CmdImg,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
		&cli.StringFlag{
			Name:    "format",
			Aliases: []string{"f"},
			Usage:   "Image output format",
			Value:   "graphviz",
		},
	},
}

var commandInit = &cli.Command{
	Name:   "init",
	Usage:  "Generate tinet config template file",
	Action: CmdInit,
}

var commandPs = &cli.Command{
	Name:   "ps",
	Usage:  "docker and netns process",
	Action: CmdPs,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "all docker and netns",
		},
	},
}

var commandPull = &cli.Command{
	Name:   "pull",
	Usage:  "Pull Node docker image from tinet config file",
	Action: CmdPull,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
	},
}

var commandReConf = &cli.Command{
	Name:   "reconf",
	Usage:  "Stop, remove, create, start and config",
	Action: CmdReConf,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Verbose",
		},
	},
}

var commandReUp = &cli.Command{
	Name:   "reup",
	Usage:  "Stop, remove, create, start",
	Action: CmdReUp,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Verbose",
		},
	},
}

var commandTest = &cli.Command{
	Name:   "test",
	Usage:  "Execute test commands from tinet config file.",
	Action: CmdTest,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
	},
}

var commandCheck = &cli.Command{
	Name:   "check",
	Usage:  "check config",
	Action: CmdCheck,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "spec.yaml",
		},
	},
}
