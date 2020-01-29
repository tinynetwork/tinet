// Package cmd tn cli tool
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tinynetwork/tn/internal/pkg/shell"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate template spec file",
	Run: func(cmd *cobra.Command, args []string) {
		tnConf, err := shell.GenerateFile()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tnConf)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
