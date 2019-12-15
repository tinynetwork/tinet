// Package cmd
package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check config",
	Run: func(cmd *cobra.Command, args []string) {
		nodes := tnconfig.Nodes
		bridges := tnconfig.Switches
		confmap := map[string]string{}

		for _, node := range nodes {
			for _, inf := range node.Interfaces {
				// fmt.Println(node.Name, " : ", inf.InfName, "->", inf.PeerNode, " : ", inf.PeerInf)
				if inf.Type == "direct" {
					host := node.Name + ":" + inf.Name
					peer := strings.Split(inf.Args, "#")
					target := peer[0] + ":" + peer[1]
					confmap[host] = target
				} else if inf.Type == "bridge" {
					host := node.Name + ":" + inf.Name
					target := inf.Args + ":" + node.Name
					confmap[host] = target
				}
			}
		}

		for _, bridge := range bridges {
			for _, inf := range bridge.Interfaces {
				host := bridge.Name + ":" + inf.Args
				target := inf.Args + ":" + inf.Name
				confmap[host] = target
			}
		}

		var matchNum int
		falseConfigMap := map[string]string{}

		for key, value := range confmap {
			if confmap[key] == value && confmap[value] == key {
				matchNum++
			} else {
				falseConfigMap[key] = value
			}
		}

		if len(confmap) == matchNum {
			log.Println("Success Check!")
		} else {
			log.Fatalf("Failed Check: %s\n", falseConfigMap)
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
