package objects_test

import (
	"os"
	"testing"

	"github.com/mattcarp12/go-gitm/gitm/files"
	"github.com/mattcarp12/go-gitm/gitm/git"
	"github.com/mattcarp12/go-gitm/gitm/objects"
)

func TestWriteTree(t *testing.T) {
	// make temp dir
	tempDir := t.TempDir()

	// chdir to temp dir
	err := os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	// initialize gitm
	git.Init(false)

	// make test tree
	testTree := map[string]interface{}{
		"a": "a",
		"b": map[string]interface{}{
			"c": "c",
			"d": "d",
		},
		"e": "e",
	}

	// write tree
	treeHash := objects.WriteTree(testTree)

	// list contents of tree
	entries, _ := os.ReadDir(files.GitmPath("objects"))

	// check that tree hash is in entries
	found := false
	for _, entry := range entries {
		t.Log(entry.Name())
		if entry.Name() == treeHash {
			found = true
		}
	}

	if !found {
		t.Error("tree hash not found in entries")
	}
}