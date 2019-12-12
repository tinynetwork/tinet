// Package cmd tn cli tool
package cmd

import (
	"fmt"
	"log"

	"github.com/ak1ra24/tn/shell"
	"github.com/spf13/cobra"
)

// reupCmd represents the reup command
var reupCmd = &cobra.Command{
	Use:   "reup",
	Short: "Stop, remove, create and start",
	Run: func(cmd *cobra.Command, args []string) {
		// stop, remove
		for _, node := range tnconfig.Nodes {
			deleteNode := shell.DeleteNode(node)
			fmt.Println(deleteNode)
		}
		for _, br := range tnconfig.Switches {
			delBrCmd := shell.DeleteSwitch(br)
			fmt.Println(delBrCmd)
		}

		// create, start
		if len(tnconfig.PreInit.Cmds) != 0 {
			shell.ExecCmd(tnconfig.PreInit.Cmds)
		}
		if len(tnconfig.PostInit.Cmds) != 0 {
			shell.ExecCmd(tnconfig.PostInit.Cmds)
		}
		for _, node := range tnconfig.Nodes {
			shell.CreateNode(node)
			shell.Mount_docker_netns(node)
		}

		if len(tnconfig.Switches) != 0 {
			for _, bridge := range tnconfig.Switches {
				shell.CreateSwitch(bridge)
			}
		}

		for _, node := range tnconfig.Nodes {
			for _, inf := range node.Interfaces {
				if inf.Type == "direct" {
					shell.N2nLink(node.Name, inf)
				} else if inf.Type == "bridge" {
					shell.S2nLink(node.Name, inf)
				} else if inf.Type == "veth" {
					shell.V2cLink(node.Name, inf)
				} else if inf.Type == "phys" {
					shell.P2cLink(node.Name, inf)
				} else {
					err := fmt.Errorf("not supported interface type: %s", inf.Type)
					log.Fatal(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(reupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
