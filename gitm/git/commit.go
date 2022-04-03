package git

import (
	"github.com/mattcarp12/go-gitm/gitm/files"
	"github.com/mattcarp12/go-gitm/gitm/index"
	"github.com/mattcarp12/go-gitm/gitm/objects"
	// "github.com/mattcarp12/go-gitm/gitm/refs"
)

// Commit creates a commit object that represents the current state of the
// index, writes the commit object to the object directory, and updates the
// HEAD reference to point to the new commit.
func Commit(message string) {
	check()

	// treeHash := writeTree()

	// var headDesc string
	// if refs.IsHeadDetached() {
	// 	headDesc = "detached HEAD"
	// } else {
	// 	headDesc = refs.HeadBranchName()
	// }



}

func writeTree() string {
	toc := index.ReadIndex().TOC()
	tree := files.NestFlatTree(toc)
	return objects.WriteTree(tree)
}