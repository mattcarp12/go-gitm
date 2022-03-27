package gitm

import (
	"log"
	"os"
	"path/filepath"
)

func objectsPath() string {
	return Files{}.GitmPath("objects")
}

func WriteObject(content []byte) string {
	hash := hashBytes(content)
	path := filepath.Join(objectsPath(), hash)
	if fileExists(path) {
		return hash
	}
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}
