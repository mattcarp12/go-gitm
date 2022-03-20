package gitm

import (
	"errors"
	"os"
	"path/filepath"
)

type Files struct{}

// gitmDir returns the absolute path of the `.gitlet` directory
// given a path inside that directory
func gitmDir(dir string) string {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return ""
	}
	// check if dir exists
	if stat, err := os.Stat(dir); !os.IsNotExist(err) && stat.IsDir() {
		// check if dir is root of repo, i.e. contains .gitm dir
		potentialGitmDir := filepath.Join(dir, ".gitm")
		if stat, err := os.Stat(potentialGitmDir); !os.IsNotExist(err) && stat.IsDir() {
			return potentialGitmDir
		} else if dir != "/" {
			return gitmDir(filepath.Clean(filepath.Join(dir, "..")))
		}
	}
	return ""
}

// GitmPath returns a string made by concatenating `path` to
// the absolute path of the `.gitlet` directory of the repo
func (f Files) GitmPath(path string) string {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	gDir := gitmDir(cwd)
	if gDir != "" {
		return filepath.Join(gDir, path)
	} else {
		return ""
	}
}

// InRepo returns true of the current working directory
// is inside a repository
func (f Files) InRepo() bool {
	return f.GitmPath("") != ""
}

// AssertInRepo panics if the current working directory
// is not inside a repository
func (f Files) AssertInRepo() {
	if !f.InRepo() {
		panic("Not in a gitm repo")
	}
}

func (f Files) RepoRoot() (string, error) {
	if !f.InRepo() {
		return "", errors.New("not in a repo")
	} else {
		return filepath.Clean(filepath.Join(f.GitmPath("."), "..")), nil
	}
}

// PathFromRepoRoot returns `path` relative to the repo root
func (f Files) PathFromRepoRoot(path string) (string, error) {
	repoRoot, err := f.RepoRoot()
	if err != nil {
		return "", err
	}
	rel, err := filepath.Rel(repoRoot, path)
	if err != nil {
		return "", err
	}
	return rel, nil
}

func (f Files) WriteFilesFromMap(tree map[string]interface{}, prefix string) error {
	for key, value := range tree {
		path := filepath.Join(prefix, key)
		switch v := value.(type) {
		case string:
			// create file with name `key` and content `value`
			err := os.WriteFile(path, []byte(v), 0777)
			if err != nil {
				return err
			}
		case map[string]interface{}:
			// create directory and recurse
			err := os.MkdirAll(path, 0777)
			if err != nil {
				return err
			}
			f.WriteFilesFromMap(v, path)
		}
	}
	return nil
}
