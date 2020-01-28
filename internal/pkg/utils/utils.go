// Package utils
package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
)

// Ask4confirm confirm yes or no
func Ask4confirm(confirmMessage string) bool {

	au := aurora.NewAurora(true)

	var s string

	message := fmt.Sprintf("%s (y/N): ", confirmMessage)
	fmt.Printf("%s", au.Bold(au.Cyan(message)))
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}

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
