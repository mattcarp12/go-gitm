package merge

import "github.com/mattcarp12/go-gitm/gitm/refs"

func IsMergeInProgress() bool {
	mergeHash := refs.Hash("MERGE_HEAD")
	return mergeHash != ""
}
