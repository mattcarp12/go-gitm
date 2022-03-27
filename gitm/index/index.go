package index

import (
	"log"
	"os"
	"strings"

	"github.com/mattcarp12/go-gitm/gitm/files"
	"github.com/mattcarp12/go-gitm/gitm/objects"
)

// Index module
// ------------

// The index maps files to hashes of their content.  When a commit is
// created, a tree is built that mirrors the content of the index.

// Index entry keys are actually a `path,stage` combination.  Stage is
// always `0`, unless the entry is about a file that is in conflict.
// See `index.writeConflict()` for more details.

type Index map[string]string

// key returns a key used in the Index object
func key(path string, stage string) string {
	return path + "," + stage
}

// keyPieces splits a key into `path` and `stage`
func keyPieces(key string) []string {
	return strings.Split(key, ",")
}

// ReadIndex reads the index as a map[string]string
func ReadIndex() Index {
	i := Index{}
	indexFilePath := files.GitmPath("index")
	if !files.FileExists(indexFilePath) {
		return i
	}
	indexBytes, err := os.ReadFile(indexFilePath)
	if err != nil {
		panic(err)
	}
	indexStr := string(indexBytes)
	indexLines := lines(indexStr)
	for _, line := range indexLines {
		lineSplit := strings.Split(line, " ")
		i[key(lineSplit[0], lineSplit[1])] = lineSplit[2]
	}
	return i
}

// TOC returns a map of the contents of the index.  The keys are
// paths, and the values are the hash of the file at that path.
func (i Index) TOC() map[string]string {
	toc := map[string]string{}
	for k := range i {
		toc[keyPieces(k)[0]] = i[k]
	}
	return toc
}

// ConflictedPaths returns a list of paths that are in conflict.
func (i Index) ConflictedPaths() []string {
	conflictedPaths := []string{}
	for k := range i {
		if keyPieces(k)[1] == "2" {
			conflictedPaths = append(conflictedPaths, keyPieces(k)[0])
		}
	}
	return conflictedPaths
}

// HasFile returns true if there is an entry for `path`
// in the index at `stage`
func (i Index) HasFile(path string, stage string) bool {
	return i[key(path, stage)] != ""
}

// IsFileInConflict returns true if the file for `path` is in conflict
func (i *Index) IsFileInConflict(path string) bool {
	return i.HasFile(path, "2")
}

// WriteRm removes the index entry for the file at `path`.
// The file will be removed from the index, even if it is in
// conflict.
func (i *Index) WriteRm(path string) {
	for _, n := range []string{"0", "1", "2", "3"} {
		delete(*i, key(path, n))
	}
}

// WriteNonConflict sets a non-conflicting index entry for the file at
// `path`.  If the file is already in conflict, it is set to no longer
// be in conflict.
func (i *Index) WriteNonConflict(path string) {
	i.WriteRm(path)
	i.WriteStage(path, "0")
}

// WriteStage adds the hashed contents of the file at `path`
// to the index at `path,stage`.
func (i *Index) WriteStage(path, stage string) {
	// read contents of file
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	(*i)[key(path, stage)] = objects.WriteObject(fileBytes)
}

// Write writes the index to .gitm/index
func (i *Index) Write() {
	indexBytes := []byte{}
	for key, value := range *i {
		keySplit := strings.Split(key, ",")
		indexBytes = append(indexBytes, []byte(keySplit[0])...)
		indexBytes = append(indexBytes, ' ')
		indexBytes = append(indexBytes, []byte(keySplit[1])...)
		indexBytes = append(indexBytes, ' ')
		indexBytes = append(indexBytes, []byte(value)...)
		indexBytes = append(indexBytes, '\n')
	}
	os.WriteFile(files.GitmPath("index"), indexBytes, 0666)
}

// updateIndex adds the contents of the file at `path` to the
// index, or removes the file from the index
func (i *Index) updateIndex(path string, add bool) {
	isOnDisk := files.PathExists(path)
	isInIndex := i.HasFile(path, "0")

	if isOnDisk && !files.FileExists(path) {
		log.Fatal("Can't add a directory to the index")
	} else if !add && !isOnDisk && isInIndex {
		// Abort if file being removed is in conflict.
		// Gitm doesn't support this
		if i.IsFileInConflict(path) {
			log.Fatal("unsupported")

			// Otherwise, remove from index
		} else {
			i.WriteRm(path)
			return
		}
	} else if !add && !isOnDisk && !isInIndex {
		// nothing to do
		return
	} else if !add && isOnDisk && !isInIndex {
		log.Fatal("file already removed from index")
	} else if isOnDisk && (add || isInIndex) {
		// either add file to index or update the index
		i.WriteNonConflict(path)
		return
	} else if add && !isOnDisk {
		log.Fatal("file not on disk...nothing to add to index")
	}
}

// Add list of files to the index
func AddFilesToIndex(files []string) {
	i := ReadIndex()
	for _, path := range files {
		i.updateIndex(path, true)
	}
	i.Write()
}

// Remove list of files from the index
func RmFilesFromIndex(files []string) {
	i := ReadIndex()
	for _, path := range files {
		i.updateIndex(path, false)
	}
	i.Write()
}

// IndexMatchingFiles returns all the paths in the index that match
// path, relative to the current directory
func IndexMatchingFiles(path string) []string {
	matches := []string{}

	toc := ReadIndex().TOC()

	searchPath := files.PathFromRepoRoot(path)

	for k := range toc {
		if strings.HasPrefix(k, searchPath) {
			// only a match if an exact match (for single file)
			// or if the next char is a path delimiter (for files in directory)
			if len(k) == len(searchPath) || k[len(searchPath)] == '/' {
				matches = append(matches, k)
			}
		}
	}

	return matches
}
