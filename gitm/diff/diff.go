package diff

import (
	"github.com/mattcarp12/go-gitm/gitm/index"
	"github.com/mattcarp12/go-gitm/gitm/objects"
	"github.com/mattcarp12/go-gitm/gitm/refs"
	"github.com/mattcarp12/go-gitm/gitm/util"
)

type FileDiff struct {
	Status   string
	Receiver string
	Base     string
	Giver    string
}

type Diff map[string]FileDiff

// WorkingCopyDiff returns a diff object that represents the difference
// between the working copy and the commit specified by `hash`.
func WorkingCopyDiff(hash string) Diff {
	commitTOC := objects.CommitTOC(hash)
	workingCopyTOC := index.WorkingCopyTOC()
	return TOCDiff(commitTOC, workingCopyTOC, nil)
}

func CommitDiff(hash1, hash2 string) Diff {
	commit1TOC := objects.CommitTOC(hash1)
	commit2TOC := objects.CommitTOC(hash2)
	return TOCDiff(commit1TOC, commit2TOC, nil)
}

// TOCDiff takes three TOC objects that map file paths to hashes of file content.
// It returns a diff between `receiver` and `giver`. `base` is the version that is
// the most recent common ancestor of `receiver` and `giver`. If `base` is not
// passed, `receiver` is used as the base. The base is only passes when getting
// the diff for a merge.
func TOCDiff(receiver, giver, base index.TOC) Diff {
	if base == nil {
		base = receiver
	}

	paths := util.Keys(receiver)
	paths = append(paths, util.Keys(giver)...)
	paths = append(paths, util.Keys(base)...)
	paths = util.Unique(paths)

	diff := Diff{}
	for _, path := range paths {
		receiverHash := receiver[path]
		giverHash := giver[path]
		baseHash := base[path]

		diff[path] = FileDiff{
			Status:   fileStatus(receiverHash, giverHash, baseHash),
			Receiver: receiverHash,
			Base:     baseHash,
			Giver:    giverHash,
		}
	}
	return diff
}

func AddedOrModifiedFiles() []string {
	// TODO - Implement this
	return []string{}
}

// ChangedFilesCommitWouldOverwrite gets a list of files
// changed in the working copy. It gets a list of the files
// that are different in the head commit and the commit for
// the passed hash. It returns a list of paths that appear
// in both lists.
func ChangedFilesCommitWouldOverwrite(hash string) []string {
	headHash := refs.Hash("HEAD")

	wcDiff := WorkingCopyDiff(headHash)
	commitDiff := CommitDiff(headHash, hash)

	return util.Intersection(util.Keys(wcDiff.NameStatus()), util.Keys(commitDiff.NameStatus()))
}
