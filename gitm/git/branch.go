package git

import (
	"fmt"
	"log"

	"github.com/mattcarp12/go-gitm/gitm/refs"
)

// Branch creates a new branch with the given name or prints existing branches
func Branch(name string) {
	check()

	// if no name is passed, list the local branches
	if name == "" {
		listLocalBranches()
		return
	}

	// `HEAD` is not pointing to a commit, so there is no commit
	// to create a branch from. Abort. This is most likely to happen
	// if the repository has no commits.
	if refs.Hash("HEAD") == "" {
		log.Println("Aborting: HEAD is not pointing to a commit.")
		return
	}

	// if the branch already exists, abort
	if refs.RefExists(refs.ToLocalRef(name)) {
		log.Printf("Aborting: branch %s already exists.", name)
		return
	}

	// create the branch
	refs.UpdateRef(refs.ToLocalRef(name), refs.Hash("HEAD"))
}

func listLocalBranches() {
	headBranch := refs.HeadBranchName()
	for branch := range refs.LocalHeads() {
		if branch == headBranch {
			fmt.Printf("* %s\n", branch)
		} else {
			fmt.Printf("  %s\n", branch)
		}
	}
}
