// Package cmd tn cli tool
package cmd

import (
	"fmt"

	"github.com/ak1ra24/tn/shell"
	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop and remove containers",
	Run: func(cmd *cobra.Command, args []string) {
		for _, node := range tnconfig.Nodes {
			deleteNode := shell.DeleteNode(node)
			fmt.Println(deleteNode)
		}
		for _, br := range tnconfig.Switches {
			delBrCmd := shell.DeleteSwitch(br)
			fmt.Println(delBrCmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
