package git

import (
	"log"

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


}
