// Package cmd tn cli tool
package cmd

import (
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print tinet config file",
	Run: func(cmd *cobra.Command, args []string) {
		pp.Print(tnconfig)
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// printCmd.PersistentFlags().StringVarP(&cfgFile, "config", "-c", "", "config file")

	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
