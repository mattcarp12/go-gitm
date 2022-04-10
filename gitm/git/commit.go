package git

import (
	"log"
	"os"
	"strings"

	"github.com/mattcarp12/go-gitm/gitm/files"
	"github.com/mattcarp12/go-gitm/gitm/index"
	"github.com/mattcarp12/go-gitm/gitm/merge"
	"github.com/mattcarp12/go-gitm/gitm/objects"
	"github.com/mattcarp12/go-gitm/gitm/refs"
)

// Commit creates a commit object that represents the current state of the
// index, writes the commit object to the object directory, and updates the
// HEAD reference to point to the new commit.
func Commit(message string) {
	check()

	// Write a tree set of tree objects that represent the current
	// state of the index
	treeHash := writeTree()

	var headDesc string
	if refs.IsHeadDetached() {
		headDesc = "detached HEAD"
	} else {
		headDesc = refs.HeadBranchName()
	}

	// Compare the hash of the tree object at the top of the tree
	// that was just written witht the hash of the tree object that the
	// `HEAD` commit points at. If they are the same, abort because
	// there is nothing new to commit.
	headRef := refs.Hash("HEAD")
	headTreeHash := objects.TreeHash(objects.Read(headRef))
	if headRef != "" && treeHash == headTreeHash {
		log.Fatal("# On " + headDesc + "\n\tnothing to commit, working directory clean")
	}

	// Abort if the repository is in the merge state and there are
	// unresolved merge conflicts
	conflictedPaths := index.ReadIndex().ConflictedPaths()
	if merge.IsMergeInProgress() && len(conflictedPaths) > 0 {
		var errMessage strings.Builder
		for _, path := range conflictedPaths {
			errMessage.WriteString("U " + path + "\n")
		}
		errMessage.WriteString("\n cannot commit because you have unmerged files\n")
		log.Fatal(errMessage.String())
	}

	// Otherwise do the commit

	// If the repository is in the merge state, use a pre-written
	// merge commit message. If the repository is not in the merge state,
	// use the message passed in with -m
	if merge.IsMergeInProgress() {
		bMessage, err := os.ReadFile(files.GitmPath("MERGE_MSG"))
		if err != nil {
			log.Fatal(err)
		}
		message = string(bMessage)
	}

	// Write new commit to the `objects` directory
	commitHash := objects.WriteCommit(treeHash, message, refs.CommitParentHashes())

	// Point `HEAD` at new commit
	refs.UpdateRef("HEAD", commitHash)

	// If `MERGE_HEAD` exists, the repository was in the merge
	// state. Remove `MERGE_HEAD` and `MERGE_MSG` to exit the
	// merge state. Report that merge is complete
	if merge.IsMergeInProgress() {
		os.Remove(files.GitmPath("MERGE_MSG"))
		refs.RM("MERGE_HEAD")
		log.Print("Merge made by the three-way strategy.")
		return
	}

	log.Print("[" + headDesc + " " + commitHash + "] " + message)
}

// writeTree takes the content of the index and stores a tree
// object that represents that content to the `objects` directory.
// It returns a hash of the tree object
func writeTree() string {
	toc := index.ReadIndex().TOC()
	tree := files.NestFlatTree(toc)
	return objects.WriteTree(tree)
}
