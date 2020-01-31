// Package utils
package utils

import (
	"fmt"
	"io"
	"os"
)

// Exists check if the directory exist
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// RemoveDuplicatesString remove duplicate string in slice
func RemoveDuplicatesString(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

// PrintCmd cmd output
func PrintCmd(w io.Writer, cmd string, verbose bool) {
	if verbose {
		fmt.Fprintln(w, cmd)
	} else {
		cmd = cmd + " > /dev/null"
		fmt.Fprintln(w, cmd)
	}
}

func PrintCmds(w io.Writer, cmds []string, verbose bool) {
	if verbose {
		for _, cmd := range cmds {
			fmt.Fprintln(w, cmd)
		}
	} else {
		for _, cmd := range cmds {
			cmd = cmd + " > /dev/null"
			fmt.Fprintln(w, cmd)
		}
	}
}
