package diff

import (
	"github.com/mattcarp12/go-gitm/gitm/index"
	"github.com/mattcarp12/go-gitm/gitm/objects"
)

type FileDiff struct {
	Status   string
	Receiver string
	Base     string
	Giver    string
}

type Diff map[string]FileDiff

func WorkingCopyDiff(hash string) Diff {
	commitTOC := objects.CommitTOC(hash)
}

func TOCDiff(receiver, giver, base index.TOC) Diff {
	// TODO - Implement this
	return Diff{}
}

func AddedOrModifiedFiles() []string {
	// TODO - Implement this
	return []string{}
}

func ChangedFilesCommitWouldOverwrite(hash string) []string {
	headHash := refs.Hash("HEAD")

}
