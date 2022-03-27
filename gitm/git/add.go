package git

import (
	"log"

	"github.com/mattcarp12/go-gitm/gitm/files"
	"github.com/mattcarp12/go-gitm/gitm/index"
)

// Add adds files that match `path` to the index.
func Add(paths []string) {
	check()

	var addedFiles []string
	for _, path := range paths {
		log.Print("Adding ", path)
		addedFiles = append(addedFiles, files.LsRecursive(path)...)
	}

	if len(addedFiles) == 0 {
		log.Fatal("nothing to add")
	}

	index.AddFilesToIndex(addedFiles)
}
