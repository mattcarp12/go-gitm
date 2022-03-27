package gitm

import (
	"log"
	"os"
)

type Git struct {
	Files Files
}

func check() {
	Files{}.AssertInRepo()
	AssertNotBare()
}

// Init initializes the current directory as a new repository
func (git Git) Init(bare bool) {
	if git.Files.InRepo() {
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
		git.Files.WriteFilesFromMap(gitmFileMap, cwd)
	} else {
		git.Files.WriteFilesFromMap(map[string]interface{}{".gitm": gitmFileMap}, cwd)
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
		addedFiles = append(addedFiles, Files{}.LsRecursive(path)...)
	}

	if len(addedFiles) == 0 {
		log.Fatal("nothing to add")
	}

	AddFilesToIndex(addedFiles)
}

func (git Git) rm(paths []string, recurse bool) {}

func (git Git) commit() {}

func (git Git) branch() {}

func (git Git) checkout() {}

func (git Git) diff() {}

func (git Git) remote() {}

func (git Git) fetch() {}

func (git Git) merge() {}

func (git Git) pull() {}

func (git Git) push() {}

func (git Git) clone() {}

func (git Git) writeTree() {}

func (git Git) updateRef() {}
