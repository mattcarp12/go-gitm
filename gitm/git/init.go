package git

import (
	"os"

	"github.com/mattcarp12/go-gitm/gitm/config"
	"github.com/mattcarp12/go-gitm/gitm/files"
)

// Init initializes the current directory as a new repository
func Init(bare bool) {
	if files.InRepo() {
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
		files.WriteFilesFromMap(gitmFileMap, cwd)
	} else {
		files.WriteFilesFromMap(map[string]interface{}{".gitm": gitmFileMap}, cwd)
	}
	// Finally, write config file
	config.WriteConfig(config.GitmConfig{Bare: bare})
}
