// Package cmd tn cli tool
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tinynetwork/tn/internal/pkg/shell"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List services",
	Run: func(cmd *cobra.Command, args []string) {
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("echo '---------------------------------------------------------------------------------'")
		fmt.Println("echo '                            Docker Status                                        '")
		fmt.Println("echo '---------------------------------------------------------------------------------'")
		dockerPsCmd := shell.DockerPs(all)
		fmt.Println(dockerPsCmd)
		fmt.Println("echo '---------------------------------------------------------------------------------'")
		fmt.Println("echo '                            IP NETNS LIST                                        '")
		fmt.Println("echo '---------------------------------------------------------------------------------'")
		netnsPsCmd := shell.NetnsPs()
		fmt.Println(netnsPsCmd)
	},
}

func init() {
	rootCmd.AddCommand(psCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	psCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
}
