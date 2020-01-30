// Package cmd tn cli tool
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tinynetwork/tn/internal/pkg/utils"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in a running container",
	Run: func(cmd *cobra.Command, args []string) {
		execCmdArgs := cmd.Flags().Args()
		execCommand := tnconfig.Exec(execCmdArgs[0], execCmdArgs[1:])
		utils.PrintCmd(os.Stdout, execCommand, verbose)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
