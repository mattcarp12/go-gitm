package git

import (
	"log"

	"github.com/mattcarp12/go-gitm/gitm"
	"github.com/mattcarp12/go-gitm/gitm/diff"
	"github.com/mattcarp12/go-gitm/gitm/files"
	"github.com/mattcarp12/go-gitm/gitm/index"
)

func Rm(path string, recurse bool) {
	check()

	filesToRm := index.IndexMatchingFiles(path)

	// Abort if no files matched `path`
	if len(filesToRm) == 0 {
		log.Fatalf("%s did not match any files", path)

		// Abort if `path` is a directory and `-r` was not passed.
	} else if files.DirExists(path) && !recurse {
		log.Fatalf("Not removing %s recursively without -r", path)

	} else {

		// Get a list of all files that are to be removed and have also
		// been changed on disk. If this list is not empty then abort.
		changesToRm := gitm.Intersection(filesToRm, diff.AddedOrModifiedFiles())
		if len(changesToRm) > 0 {
			errMsg := "the following files have changes:\n"
			for _, changedFile := range changesToRm {
				errMsg += changedFile + "\n"
			}
			log.Fatal(errMsg)

			// Otherwise, remove the files that match `path`. Delete them
			// from disk and remove from the index.
		} else {
			files.DeleteFiles(filesToRm)
			index.RmFilesFromIndex(filesToRm)
		}
	}
}
