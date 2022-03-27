package gitm

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGitmDir(t *testing.T) {
	// Make temp dir
	tmp := t.TempDir()

	// Add .gitm to dir
	os.MkdirAll(filepath.Join(tmp, ".gitm"), os.ModeDir)

	// Change working directory
	os.Chdir(tmp)

	gDir := gitmDir(tmp)
	expected := filepath.Join(tmp, ".gitm")
	if gDir != expected {
		t.Errorf("%s != %s", gDir, expected)
	}

	// Check for failure
	tmp = t.TempDir()
	os.Chdir(tmp)
	gDir = gitmDir(tmp)
	if gDir != "" {
		t.Errorf("%s is not empty", gDir)
	}
}

func TestRepoRoot(t *testing.T) {
	// Make temp dir
	tmp := t.TempDir()

	// Add .gitm to dir
	os.MkdirAll(filepath.Join(tmp, ".gitm"), os.ModeDir)

	// Change working directory
	os.Chdir(tmp)

	rootDir := RepoRoot()
	if rootDir != tmp {
		t.Errorf("%s != %s", rootDir, tmp)
	}

}

func TestWriteFilesFromMap(t *testing.T) {
	fileMap := map[string]interface{}{
		"foo": "bar",
		"bar": map[string]interface{}{
			"foobar": "foobar",
		},
	}
	tmp := t.TempDir()
	WriteFilesFromMap(fileMap, tmp)
	
	// Check if files were created
	if stat, err := os.Stat(filepath.Join(tmp, "foo")); os.IsNotExist(err) || !stat.Mode().IsRegular() {
		t.Error(err)
	}
	if stat, err := os.Stat(filepath.Join(tmp, "bar")); os.IsNotExist(err) || !stat.Mode().IsDir() {
		t.Error(err)
	}
	if stat, err := os.Stat(filepath.Join(tmp, "bar", "foobar")); os.IsNotExist(err) || !stat.Mode().IsRegular() {
		t.Error(err)
	}
	
	// Check file contents
	contents, _ := os.ReadFile(filepath.Join(tmp, "foo"))
	if string(contents) != "bar" {
		t.Error("Incorrect contents")
	}
}
