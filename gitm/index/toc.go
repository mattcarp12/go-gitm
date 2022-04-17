package index

import (
	"log"
	"os"

	"github.com/mattcarp12/go-gitm/gitm/files"
	"github.com/mattcarp12/go-gitm/gitm/objects"
	"github.com/mattcarp12/go-gitm/gitm/util"
)

type TOC map[string]string

// TOC returns a map of the contents of the index.  The keys are
// paths, and the values are the hash of the file at that path.
func (i Index) TOC() TOC {
	toc := TOC{}
	for k := range i {
		toc[keyPieces(k)[0]] = i[k]
	}
	return toc
}

func WorkingCopyTOC() TOC {
	paths := util.Keys(ReadIndex())

	toc := TOC{}
	for _, path := range paths {
		path := keyPieces(path)[0]

		// read file contents if it exists
		if files.FileExists(path) {
			fileContents, err := os.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}
			toc[path] = objects.Hash(fileContents)
		}
	}
	return toc
}

func TocToIndex(toc TOC) Index {
	index := Index{}
	for path, hash := range toc {
		index[key(path, "0")] = hash
	}
	return index
}