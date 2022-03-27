package gitm

import (
	"log"
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
	// make sure dir exists
	if stat, err := os.Stat(dir); !os.IsNotExist(err) && stat.IsDir() {
		potentialConfigFile := filepath.Join(dir, "config")
		potentialGitmDir := filepath.Join(dir, ".gitm")
		// check if dir contains config file
		if stat, err := os.Stat(potentialConfigFile); !os.IsNotExist(err) && stat.Mode().IsRegular() {
			return dir
			// check if dir is root of repo, i.e. contains .gitm dir
		} else if stat, err := os.Stat(potentialGitmDir); !os.IsNotExist(err) && stat.IsDir() {
			return potentialGitmDir
		} else if dir != "/" {
			// If above checks failed, recurse in parent directory (until reach root)
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

func (f Files) RepoRoot() string {
	f.AssertInRepo()
	return filepath.Clean(filepath.Join(f.GitmPath("."), ".."))
}

// PathFromRepoRoot returns `path` relative to the repo root
func (f Files) PathFromRepoRoot(path string) string {
	repoRoot := f.RepoRoot()
	rel, err := filepath.Rel(repoRoot, path)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return rel
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

func (f Files) EqualsGitm(path string) bool {
	return f.GitmPath(path) == path
}

func (f Files) LsRecursive(path string) []string {
	if !pathExists(path) {
		return []string{}
	} else if absPath, _ := filepath.Abs(path); absPath == gitmDir(path) {
		return []string{}
	} else if fileExists(path) {
		log.Print("Found file: ", path)
		return []string{path}
	} else if dirExists(path) {
		log.Print("Found directory: ", path)
		dirEnt, err := os.ReadDir(path)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		files := []string{}
		for _, file := range dirEnt {
			files = append(files, f.LsRecursive(filepath.Join(path, file.Name()))...)
		}
		return files
	}
	return []string{}
}
