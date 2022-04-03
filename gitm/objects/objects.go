package objects

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattcarp12/go-gitm/gitm/files"
)

// Objects module
// -----------

// Objects are files in the `.gitm/objects` directory.
// - A blob object stores the contents of a file.
// - A tree object stores a list of files and directories in a directory.
//   Entries in the list for files point to blob objects. Entries
//   in the list for directories point to tree objects.
// - A commit object stores a tree object and a list of parent commit
//   objects, plus a message.
// - A tag object stores a commit object and a name.
// - A branch object stores a commit object and a name.
// - A remote object stores a URL.


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

// WriteTree stores a graph of tree objects that represent the
// content currently in the index, and returns
// a hash of the root tree object.
func WriteTree(tree map[string]interface{}) string {
	var treeObject strings.Builder
	for key, value := range tree {
		switch v := value.(type) {
		case string:
			treeObject.WriteString("blob ")
			treeObject.WriteString(v)
		case map[string]interface{}:
			treeObject.WriteString("tree ")
			treeObject.WriteString(WriteTree(v))
		}
		treeObject.WriteString(" " + key)
		treeObject.WriteString("\n")
	}
	return WriteObject([]byte(treeObject.String()))
}

// Exists returns true if the object with the given hash exists
func Exists(hash string) bool {
	path := filepath.Join(objectsPath(), hash)
	return files.FileExists(path)
}

func Read(hash string) string {
	path := filepath.Join(objectsPath(), hash)
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func Type(hash string) string {
	content := Read(hash)
	return string(content[:8])
}