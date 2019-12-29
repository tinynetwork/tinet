// Package cmd tn cli tool
package cmd

import (
	"fmt"
	"strings"

	"github.com/ak1ra24/tn/internal/pkg/shell"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Execute tests",
	Run: func(cmd *cobra.Command, args []string) {
		tnTestCmds := shell.TnTestCmdExec(tnconfig.Test)
		fmt.Println(strings.Join(tnTestCmds, "\n"))

	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
