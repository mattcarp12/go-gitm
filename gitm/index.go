package gitm

import (
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

func (i Index) HasFile(path string, stage string) bool {
	return i[key(path, stage)] != ""
}

func (i Index) IsFileInConflict(path string) bool {
	return i.HasFile(path, "2")
}

func (i Index) WriteRm(path string) {}

func (i Index) WriteNonConflict(path string, content string) {}

func (i Index) Write() {
	indexBytes := []byte{}
	for key, value := range i {
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