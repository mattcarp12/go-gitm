package git

import (
	"log"
	"os"

	"github.com/mattcarp12/go-gitm/gitm/diff"
	"github.com/mattcarp12/go-gitm/gitm/files"
	"github.com/mattcarp12/go-gitm/gitm/index"
	"github.com/mattcarp12/go-gitm/gitm/objects"
	"github.com/mattcarp12/go-gitm/gitm/refs"
)

// checkout() changes the index, working copy and `HEAD` to
// reflect the content of `ref`. `ref` can be a branch name, a commit
// hash, or a tag name.
func Checkout(ref string) {
	check()

	// get the toHash of the ref
	toHash := refs.Hash(ref)

	// if the ref cant be found, abort
	if toHash == "" {
		log.Fatalf("Aborting: ref %s not found.", ref)
	}

	// abort if the hash points to a non-commit object
	if objects.Type(objects.Read(toHash)) != "commit" {
		log.Fatalf("Aborting: ref %s does not point to a commit.", ref)
	}

	// Abort if `ref` is the same as the branch that `HEAD` is pointing to
	if refs.HeadBranchName() == ref {
		log.Printf("Already on %s", ref)
		return
	}

	// Get a list of files changed in the working copy. Get a list
	// of the files that are different in the head commit and the
	// commit to check out. If any files appear in both lists, abort.
	// This is to prevent the user from overwriting files that have
	// been changed since the last commit.
	changedFiles := diff.ChangedFilesCommitWouldOverwrite(toHash)
	if len(changedFiles) > 0 {
		log.Println("Aborting: the following files have been changed since the last commit: ")
		for _, file := range changedFiles {
			log.Printf("\t%s", file)
		}
		return
	}

	// Perform the checkout
	err := os.Chdir(files.RepoRoot())
	if err != nil {
		log.Fatalf("Aborting: %s", err)
	}

	// Get the list of differences between the head commit and the
	// commit to checkout. Write them to the working copy.
	diff.CommitDiff(refs.Hash("HEAD"), toHash).WriteWorkingCopy()

	// If the ref is in the objects directory, it must be a hash,
	// so this checkout is detaching the head
	isDetachingHead := objects.Exists(ref)

	// Write the commit being checked out the `HEAD`. If the head
	// is being detached, the commit hash is written directly to
	// `HEAD`. Otherwise, the branch being checked out is written
	// to `HEAD`.
	if isDetachingHead {
		refs.WriteRef("HEAD", toHash)
	} else {
		refs.WriteRef("HEAD", "refs: "+refs.ToLocalRef(ref))
	}

	// Set the index to the contents of the commit being checked out
	idx := index.TocToIndex(objects.CommitTOC(toHash))
	(&idx).Write()

	// Report the result of the checkout
	if isDetachingHead {
		log.Printf("Note: checking out commit %s\nYou are in detached HEAD state.", toHash)
	} else {
		log.Printf("Switched to branch %s", ref)
	}
}
