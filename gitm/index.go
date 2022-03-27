package gitm

import (
	"log"
	"os"
	"strings"
)

// Index module
// ------------

// The index maps files to hashes of their content.  When a commit is
// created, a tree is built that mirrors the content of the index.

// Index entry keys are actually a `path,stage` combination.  Stage is
// always `0`, unless the entry is about a file that is in conflict.
// See `index.writeConflict()` for more details.

type Index map[string]string

func key(path string, stage string) string {
	return path + "," + stage
}

func keyPieces(key string) []string {
	return strings.Split(key, ",")
}

func ReadIndex() Index {
	f := Files{}
	i := Index{}
	indexFilePath := f.GitmPath("index")
	if !fileExists(indexFilePath) {
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

func (i Index) HasFile(path string, stage string) bool {
	return i[key(path, stage)] != ""
}

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

// WriteNonConflict sets a non-conflicting index entry fot the file at
// `path`.  If the file is already in conflict, it is set to no longer
// be in conflict.
func (i *Index) WriteNonConflict(path string) {
	i.WriteRm(path)
	i.WriteStage(path, "0")
}

func (i *Index) WriteStage(path, stage string) {
	// read contents of file
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	(*i)[key(path, stage)] = WriteObject(fileBytes)
}

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
	os.WriteFile(Files{}.GitmPath("index"), indexBytes, 0666)
}


// updateIndex adds the contents of the file at `path` to the
// index, or removes the file from the index
func (i *Index) updateIndex(path string, add bool) {

	isOnDisk := pathExists(path)
	isInIndex := i.HasFile(path, "0")

	if isOnDisk && !fileExists(path) {
		log.Fatal("Can't add a directory to the index")
	} else if !add && !isOnDisk && isInIndex {
		// Abort if file being removed is in conflict.
		// Gitm doesn't support this
		if (i.IsFileInConflict(path)) {
			log.Fatal("unsupported")
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

func AddFilesToIndex(files []string) {
	i := ReadIndex()
	for _, path := range files {
		i.updateIndex(path, true)
	}
	i.Write()
}