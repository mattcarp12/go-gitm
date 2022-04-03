package refs_test

import (
	"os"
	"testing"

	"github.com/mattcarp12/go-gitm/gitm/git"
	"github.com/mattcarp12/go-gitm/gitm/refs"
)

func testInit(t *testing.T) string {
	// Create temp working directory in which to run tests
	tempDir := t.TempDir()

	// Change working directory to temp directory
	err := os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	// Initialize gitm
	git.Init(false)

	return tempDir
}

func TestHeadBranchName(t *testing.T) {
	testInit(t)

	headBranch := refs.HeadBranchName()
	t.Logf("HEAD branch: %s", headBranch)
	if headBranch != "master" {
		t.Fatal("HEAD should be master")
	}
}