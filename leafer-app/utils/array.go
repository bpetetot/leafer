package utils

import (
	"os"
	"sort"
)

// Contains check if a string is in an array of strings
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// SortFiles will the files following their natural order of names
func SortFiles(files []os.FileInfo) []os.FileInfo {
	sort.Slice(files[:], func(i, j int) bool {
		return NaturalLess(files[i].Name(), files[j].Name())
	})
	return files
}
