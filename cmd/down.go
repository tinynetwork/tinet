// Package cmd tn cli tool
package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tinynetwork/tn/internal/pkg/utils"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop and remove containers",
	Run: func(cmd *cobra.Command, args []string) {
		for _, node := range tnconfig.Nodes {
			deleteNode := node.DeleteNode()
			utils.PrintCmd(os.Stdout, strings.Join(deleteNode, "\n"), verbose)
		}
		for _, br := range tnconfig.Switches {
			delBrCmd := br.DeleteSwitch()
			utils.PrintCmd(os.Stdout, delBrCmd, verbose)
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
