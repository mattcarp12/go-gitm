package gitm

import (
	"log"
	"os"
)

type Git struct{}

func check() {
	AssertInRepo()
	AssertNotBare()
}

// Init initializes the current directory as a new repository
func (git Git) Init(bare bool) {
	if InRepo() {
		return
	}

	// Map that mirrors the basic Git directory structure
	gitmFileMap := map[string]interface{}{
		"HEAD":    "ref: refs/heads/master\n",
		"config":  "",
		"objects": map[string]interface{}{}, // empty directory
		"refs": map[string]interface{}{
			"heads": map[string]interface{}{},
		},
	}

	// Write files to cwd
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if bare {
		WriteFilesFromMap(gitmFileMap, cwd)
	} else {
		WriteFilesFromMap(map[string]interface{}{".gitm": gitmFileMap}, cwd)
	}
	// Finally, write config file
	WriteConfig(GitmConfig{Bare: bare})
}

// Add adds files that match `path` to the index.
func (git Git) Add(paths []string) {
	check()

	var addedFiles []string
	for _, path := range paths {
		log.Print("Adding ", path)
		addedFiles = append(addedFiles, LsRecursive(path)...)
	}

	if len(addedFiles) == 0 {
		log.Fatal("nothing to add")
	}

	AddFilesToIndex(addedFiles)
}

func (git Git) Rm(path string, recurse bool) {
	check()

	filesToRm := IndexMatchingFiles(path)

	// Abort if no files matched `path`
	if len(filesToRm) == 0 {
		log.Fatalf("%s did not match any files", path)

		// Abort if `path` is a directory and `-r` was not passed.
	} else if dirExists(path) && !recurse {
		log.Fatalf("Not removing %s recursively without -r", path)

	} else {

		// Get a list of all files that are to be removed and have also
		// been changed on disk. If this list is not empty then abort.
		changesToRm := intersection(filesToRm, AddedOrModifiedFiles())
		if len(changesToRm) > 0 {
			errMsg := "the following files have changes:\n"
			for _, changedFile := range changesToRm {
				errMsg += changedFile + "\n"
			}
			log.Fatal(errMsg)
		
		// Otherwise, remove the files that match `path`. Delete them
		// from disk and remove from the index.
		} else {
			DeleteFiles(filesToRm)
			RmFilesFromIndex(filesToRm)
		}
	}
}

func (git Git) Commit() {}

func (git Git) Branch() {}

func (git Git) Checkout() {}

func (git Git) Diff() {}

func (git Git) Remote() {}

func (git Git) Fetch() {}

func (git Git) Merge() {}

func (git Git) Pull() {}

func (git Git) Push() {}

func (git Git) Clone() {}

func (git Git) WriteTree() {}

func (git Git) UpdateRef() {}
