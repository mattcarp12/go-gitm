package refs

// Refs module
// -----------

// Refs are names for commit hashes. The ref is the name of a file.
// Some refs represent local branches, like `refs/heads/master`.
// Some refs represent remote branches, like `refs/remotes/origin/master`.
// Some refs represent tags, like `refs/tags/v1.0.0`. Some represent important
// states of the repository, like `refs/heads/HEAD`.
// Ref files contain either a hash or a ref to another ref.

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/mattcarp12/go-gitm/gitm/files"
	"github.com/mattcarp12/go-gitm/gitm/objects"
)

func isRef(ref string) bool { return false }

// TerminalRef returns ref to the most specific ref
func TerminalRef(ref string) string {
	// If ref is HEAD, return the ref to the branch
	if ref == "HEAD" && !IsHeadDetached() {
		headPath := files.GitmPath("HEAD")
		headContents, err := os.ReadFile(headPath)
		if err != nil {
			log.Fatal(err)
		}
		return regexp.MustCompile(`ref: (refs/heads/.+)`).FindStringSubmatch(string(headContents))[1]

		// If ref is qualified, return the ref
	} else if isRef(ref) {
		return ref

		// If ref is a branch, return the ref to the branch
	} else {
		return ToLocalRef(ref)
	}
}

// Hash returns the hash that refOrHash points to
func Hash(refOrHash string) string {
	if objects.Exists(refOrHash) {
		// refOrHash is a hash
		return refOrHash
	} else {
		terminalRef := TerminalRef(refOrHash)
		if terminalRef == "FETCH_HEAD" {
			log.Fatal("NOT IMPLEMENTED")
		} else if RefExists(terminalRef) {
			refContents, err := os.ReadFile(files.GitmPath(terminalRef))
			if err != nil {
				log.Fatal(err)
			}
			return string(refContents)
		} else {
			log.Print("ref does not exist: " + terminalRef)
			return ""
		}
	}
	return ""
}

// IsHeadDetached returns true if `HEAD` contains a commit
// hash, rather than the ref to a branch
func IsHeadDetached() bool {
	headPath := files.GitmPath("HEAD")
	headContents, err := os.ReadFile(headPath)
	if err != nil {
		log.Fatal(err)
	}
	return !strings.Contains(string(headContents), "ref")
}

func HeadBranchName() string {
	if IsHeadDetached() {
		log.Fatal("HEAD is detached")
	}
	headPath := files.GitmPath("HEAD")
	headContents, err := os.ReadFile(headPath)
	if err != nil {
		log.Fatal(err)
	}
	return regexp.MustCompile(`refs/heads/(.+)`).FindStringSubmatch(string(headContents))[1]
}

func isCheckedOut(branch string) bool {
	return true
}

func ToLocalRef(name string) string {
	return "refs/heads/" + name
}

func toRemoteRef(remote, name string) string {
	return "refs/remotes/" + remote + "/" + name
}

func writeRef(ref, content string) {
	if isRef(ref) {
		os.WriteFile(files.GitmPath(ref), []byte(content), 0644)
	}
}

func rm() {}

// RefExists returns true if the qualified ref exists
func RefExists(qualifiedRef string) bool {
	return isRef(qualifiedRef) && files.FileExists(files.GitmPath(qualifiedRef))
}

// LocalHeads returns a map of local refs to their hashes
func LocalHeads() map[string]string {
	headsDir := files.GitmPath("refs/heads")
	if !files.DirExists(headsDir) {
		log.Fatal("refs/heads does not exist")
	}
	heads, err := os.ReadDir(headsDir)
	if err != nil {
		log.Fatal(err)
	}
	refs := make(map[string]string)
	for _, head := range heads {
		refs[head.Name()] = Hash(head.Name())
	}
	return refs
}

// UpdateRef gets the hash of the commit that `refToUpdateTo` points to,
// and sets `refToUpdate` to point to that hash
func UpdateRef(refToUpdate, refToUpdateTo string) {
	files.AssertInRepo()

	// get the hash that `refToUpdateTo` points to
	hash := Hash(refToUpdateTo)

	// Abort if `refToUpdateTo` is does not point to a valid object
	if !objects.Exists(hash) {
		log.Fatal("refToUpdateTo does not point to a valid object")
	}

	// Abort if `refToUpdate` is not a valid ref
	if !isRef(refToUpdate) {
		log.Fatal("refToUpdate is not a valid ref")
	}

	// Abort if `hash` points to an object in the `objects` directory
	// that is not a commit.
	if objects.Type(objects.Read(hash)) != "commit" {
		branch := TerminalRef(refToUpdate)
		log.Fatal(branch + " cannot refer to a non-commit object " + hash)
	}

	// Otherwise, set the contents of the file at `refToUpdate` to `hash`
	writeRef(refToUpdate, hash)
}
