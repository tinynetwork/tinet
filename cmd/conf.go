// Package cmd tn cli tool
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "Execute config-cmd in a running container",
	Run: func(cmd *cobra.Command, args []string) {
		nodeinfo := map[string]string{}
		for _, node := range tnconfig.Nodes {
			nodeinfo[node.Name] = node.Type
		}

		for _, nodeConfig := range tnconfig.NodeConfigs {
			execConfCmds := nodeConfig.ExecConf(nodeinfo[nodeConfig.Name])
			for _, execConfCmd := range execConfCmds {
				fmt.Println(execConfCmd)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(confCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// confCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// confCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
