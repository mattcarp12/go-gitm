package gitm

import (
	"fmt"
	"os"
)

type Git struct{}

func (git Git) Init(bare bool) {
	f := Files{}
	if f.InRepo() {
		return
	}

	// Map that mirrors the basic Git directory structure
	gitmFileMap := map[string]interface{}{
		"HEAD": "ref: refs/heads/master\n",
		"objects": map[string]interface{}{}, // empty directory
		"refs": map[string]interface{}{
			"heads": map[string]interface{}{},
		},
		// TODO -- Write config file
	}

	// Write files to cwd
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bare? %t\n", bare)
	if bare {
		f.WriteFilesFromMap(gitmFileMap, cwd)
	} else {
		f.WriteFilesFromMap(map[string]interface{}{".gitm": gitmFileMap}, cwd)
	}
}

func (git Git) add() {}

func (git Git) rm() {}

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

func (git Git) updateIndex() {}

func (git Git) writeTree() {}

func (git Git) updateRef() {}
