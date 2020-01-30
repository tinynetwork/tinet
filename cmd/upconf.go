// Package cmd tn cli tool
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tinynetwork/tinet/internal/pkg/shell"
	"github.com/tinynetwork/tinet/internal/pkg/utils"
)

// upconfCmd represents the upconf command
var upconfCmd = &cobra.Command{
	Use:   "upconf",
	Short: "Create, start and config",
	Run: func(cmd *cobra.Command, args []string) {
		// up
		if len(tnconfig.PreCmd) != 0 {
			for _, preCmds := range tnconfig.PreCmd {
				preExecCmds := shell.ExecCmd(preCmds.Cmds)
				utils.PrintCmd(os.Stdout, strings.Join(preExecCmds, "\n"), verbose)
			}
		}
		if len(tnconfig.PreInit) != 0 {
			for _, preInitCmds := range tnconfig.PreInit {
				preExecInitCmds := shell.ExecCmd(preInitCmds.Cmds)
				utils.PrintCmd(os.Stdout, strings.Join(preExecInitCmds, "\n"), verbose)
			}
		}
		for _, node := range tnconfig.Nodes {
			createNodeCmds := node.CreateNode()
			utils.PrintCmd(os.Stdout, strings.Join(createNodeCmds, "\n"), verbose)

			if node.Type != "netns" {
				mountDockerNetnsCmds := node.Mount_docker_netns()
				utils.PrintCmd(os.Stdout, strings.Join(mountDockerNetnsCmds, "\n"), verbose)
			}
		}

		if len(tnconfig.Switches) != 0 {
			for _, bridge := range tnconfig.Switches {
				createSwitchCmds := bridge.CreateSwitch()
				utils.PrintCmd(os.Stdout, strings.Join(createSwitchCmds, "\n"), verbose)
			}
		}

		for _, node := range tnconfig.Nodes {
			for _, inf := range node.Interfaces {
				if inf.Type == "direct" {
					n2nLinkCmds := inf.N2nLink(node.Name)
					utils.PrintCmd(os.Stdout, strings.Join(n2nLinkCmds, "\n"), verbose)
				} else if inf.Type == "bridge" {
					s2nLinkCmds := inf.S2nLink(node.Name)
					utils.PrintCmd(os.Stdout, strings.Join(s2nLinkCmds, "\n"), verbose)
				} else if inf.Type == "veth" {
					v2cLinkCmds := inf.V2cLink(node.Name)
					utils.PrintCmd(os.Stdout, strings.Join(v2cLinkCmds, "\n"), verbose)
				} else if inf.Type == "phys" {
					p2cLinkCmds := inf.P2cLink(node.Name)
					utils.PrintCmd(os.Stdout, strings.Join(p2cLinkCmds, "\n"), verbose)
				} else {
					err := fmt.Errorf("not supported interface type: %s", inf.Type)
					log.Fatal(err)
				}
			}
		}

		if len(tnconfig.PostInit) != 0 {
			for _, postInitCmds := range tnconfig.PostInit {
				postExecInitCmds := shell.ExecCmd(postInitCmds.Cmds)
				utils.PrintCmd(os.Stdout, strings.Join(postExecInitCmds, "\n"), verbose)
			}
		}

		// conf
		nodeinfo := map[string]string{}
		for _, node := range tnconfig.Nodes {
			nodeinfo[node.Name] = node.Type
		}

		for _, nodeConfig := range tnconfig.NodeConfigs {
			execConfCmds := nodeConfig.ExecConf(nodeinfo[nodeConfig.Name])
			for _, execConfCmd := range execConfCmds {
				utils.PrintCmd(os.Stdout, execConfCmd, verbose)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(upconfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upconfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upconfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
