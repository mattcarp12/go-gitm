package objects

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mattcarp12/go-gitm/gitm/files"
)

func objectsPath() string {
	return files.GitmPath("objects")
}

// WriteObject writes content to the objects database,
// and returns the hash of the contents
func WriteObject(content []byte) string {
	hash := hashBytes(content)
	path := filepath.Join(objectsPath(), hash)
	if files.FileExists(path) {
		return hash
	}
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}
