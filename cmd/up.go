// Package cmd tn cli tool
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/ak1ra24/tn/shell"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Create and start containers",
	Run: func(cmd *cobra.Command, args []string) {
		if len(tnconfig.PreCmd) != 0 {
			for _, preCmds := range tnconfig.PreCmd {
				preExecCmds := shell.ExecCmd(preCmds.Cmds)
				fmt.Println(strings.Join(preExecCmds, "\n"))
			}
		}
		if len(tnconfig.PreInit) != 0 {
			for _, preInitCmds := range tnconfig.PreInit {
				preExecInitCmds := shell.ExecCmd(preInitCmds.Cmds)
				fmt.Println(strings.Join(preExecInitCmds, "\n"))
			}
		}
		for _, node := range tnconfig.Nodes {
			createNodeCmds := shell.CreateNode(node)
			fmt.Println(strings.Join(createNodeCmds, "\n"))

			if node.Type != "netns" {
				mountDockerNetnsCmds := shell.Mount_docker_netns(node)
				fmt.Println(strings.Join(mountDockerNetnsCmds, "\n"))
			}
		}

		if len(tnconfig.Switches) != 0 {
			for _, bridge := range tnconfig.Switches {
				createSwitchCmds := shell.CreateSwitch(bridge)
				fmt.Println(strings.Join(createSwitchCmds, "\n"))
			}
		}

		for _, node := range tnconfig.Nodes {
			for _, inf := range node.Interfaces {
				if inf.Type == "direct" {
					n2nLinkCmds := shell.N2nLink(node.Name, inf)
					fmt.Println(strings.Join(n2nLinkCmds, "\n"))
				} else if inf.Type == "bridge" {
					s2nLinkCmd := shell.S2nLink(node.Name, inf)
					fmt.Println(strings.Join(s2nLinkCmd, "\n"))
				} else if inf.Type == "veth" {
					v2cLinkCmds := shell.V2cLink(node.Name, inf)
					fmt.Println(strings.Join(v2cLinkCmds, "\n"))
				} else if inf.Type == "phys" {
					p2cLinkCmds := shell.P2cLink(node.Name, inf)
					fmt.Println(strings.Join(p2cLinkCmds, "\n"))
				} else {
					err := fmt.Errorf("not supported interface type: %s", inf.Type)
					log.Fatal(err)
				}
			}
		}

		if len(tnconfig.PostInit) != 0 {
			for _, postInitCmds := range tnconfig.PostInit {
				postExecInitCmds := shell.ExecCmd(postInitCmds.Cmds)
				fmt.Println(strings.Join(postExecInitCmds, "\n"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
