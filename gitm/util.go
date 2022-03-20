package gitm

import (
	"os"
	"strings"
)

// lines takes a string, splits on newlines, and returns an
// array of the lines that are not empty
func lines(str string) []string {
	lines1 := strings.Split(str, "\n")
	lines2 := []string{}
	for _, line := range lines1 {
		if line != "" {
			lines2 = append(lines2, line)
		}
	}
	return lines2
}

// fileExists returns true if path exists and is a regular file
func fileExists(path string) bool {
	if stat, err := os.Stat(path); !os.IsNotExist(err) && stat.Mode().IsRegular() {
		return true
	}
	return false
}

// dirExists returns true if path exists and is a directory
func dirExists(path string) bool {
	if stat, err := os.Stat(path); !os.IsNotExist(err) && stat.IsDir() {
		return true
	}
	return false
}
