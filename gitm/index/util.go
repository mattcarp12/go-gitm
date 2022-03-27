package index

import "strings"

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
