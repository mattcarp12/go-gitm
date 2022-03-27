package gitm

import (
	"crypto/sha1"
	"encoding/hex"
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

func pathExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
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

func hashStr(str string) string {
	arr := sha1.Sum([]byte(str))
	return hex.EncodeToString(arr[:])
}

func hashBytes(bytes []byte) string {
	arr := sha1.Sum(bytes)
	return hex.EncodeToString(arr[:])
}

func intersection(a, b []string) []string {
	res := []string{}

	for _, vala := range a {
		for _, valb := range b {
			if vala == valb {
				res = append(res, vala)
			}
		}
	}

	return res
}