// Package cmd
package cmd

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/spf13/cobra"
)

// imgCmd represents the img command
var imgCmd = &cobra.Command{
	Use:   "img",
	Short: "Generate topology png file by graphviz",
	Run: func(cmd *cobra.Command, args []string) {
		g := dot.NewGraph(dot.Directed)
		for _, tnnode := range tnconfig.Nodes {
			nodeName := tnnode.Name
			fromNode := g.Node(nodeName)
			fromNode.Label(nodeName)
			ifaceInfos := tnnode.Interfaces
			for _, ifaceInfo := range ifaceInfos {
				var argsName string
				if ifaceInfo.Type == "direct" {
					argsName = strings.Split(ifaceInfo.Args, "#")[0]
					toNode := g.Node(argsName)
					findEdges := g.FindEdges(toNode, fromNode)
					if len(findEdges) == 0 {
						newEdge := g.Edge(fromNode, toNode)
						newEdge.Attr("arrowhead", "none")
						newEdge.Attr("labelfloat", "true")
						newEdge.Attr("headlabel", strings.Split(ifaceInfo.Args, "#")[1])
						newEdge.Attr("taillabel", ifaceInfo.Name)
						newEdge.Attr("fontsize", "8")
					}
				} else if ifaceInfo.Type == "bridge" {
					argsName = ifaceInfo.Args
					toNode := g.Node(argsName)
					findEdges := g.FindEdges(toNode, fromNode)
					if len(findEdges) == 0 {
						newEdge := g.Edge(fromNode, toNode)
						newEdge.Attr("arrowhead", "none")
						newEdge.Attr("labelfloat", "true")
						newEdge.Attr("headlabel", argsName)
						newEdge.Attr("taillabel", ifaceInfo.Name)
						newEdge.Attr("fontsize", "8")
					}
				}
			}
		}
		fmt.Println(g.String())
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
