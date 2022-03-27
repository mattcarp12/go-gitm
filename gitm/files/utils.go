package files

import "os"

func PathExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

// FileExists returns true if path exists and is a regular file
func FileExists(path string) bool {
	if stat, err := os.Stat(path); !os.IsNotExist(err) && stat.Mode().IsRegular() {
		return true
	}
	return false
}

// DirExists returns true if path exists and is a directory
func DirExists(path string) bool {
	if stat, err := os.Stat(path); !os.IsNotExist(err) && stat.IsDir() {
		return true
	}
	return false
}
