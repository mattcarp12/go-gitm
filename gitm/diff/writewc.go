package diff

import (
	"os"

	"github.com/mattcarp12/go-gitm/gitm/objects"
)

func composeConflict(receiverHash, giverHash string) string {
	return "<<<<<<\n" + objects.Read(receiverHash) +
		"\n======\n" + objects.Read(giverHash) +
		"\n>>>>>>\n"
}

func (d Diff) WriteWorkingCopy() {
	for _, path := range d.keys() {
		fileDiff := d[path]
		switch fileDiff.Status {
		case Add:
			if fileDiff.Giver == "" {
				os.WriteFile(path, []byte(objects.Read(fileDiff.Receiver)), 0644)
			} else {
				os.WriteFile(path, []byte(objects.Read(fileDiff.Giver)), 0644)
			}
		case Delete:
			os.Remove(path)
		case Modify:
			os.WriteFile(path, []byte(objects.Read(fileDiff.Giver)), 0644)
		case Conflict:
			os.WriteFile(path, []byte(composeConflict(fileDiff.Receiver, fileDiff.Giver)), 0644)
		}
	}
}
