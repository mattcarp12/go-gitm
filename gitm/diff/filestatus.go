package diff

const (
	Add      = "A"
	Modify   = "M"
	Delete   = "D"
	Same     = "SAME"
	Conflict = "Conflict"
)

// fileStatus takes three
func fileStatus(receiver, giver, base string) string {
	receiverPresent := receiver != ""
	giverPresent := giver != ""
	basePresent := base != ""

	if receiverPresent && giverPresent && receiver != giver {
		if receiver != base && giver != base {
			return Conflict
		} else {
			return Modify
		}
	} else if receiver == giver {
		return Same
	} else if !receiverPresent && !basePresent && giverPresent {
		return Add
	} else if receiverPresent && !basePresent && !giverPresent {
		return Add
	} else if receiverPresent && basePresent && !giverPresent {
		return Delete
	} else if !receiverPresent && basePresent && giverPresent {
		return Delete
	}
	return ""
}

func (d Diff) NameStatus() map[string]string {
	nameStatus := map[string]string{}
	for path, fileDiff := range d {
		if fileDiff.Status == Same {
			continue
		}
		nameStatus[path] = fileDiff.Status
	}
	return nameStatus
}
