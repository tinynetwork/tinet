// Package cmd tn cli tool
package cmd

import (
	"log"

	"github.com/ak1ra24/tn/shell"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generate a Docker bundle from the spec file",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal(shell.BuildCmd(tnconfig.Nodes))
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
