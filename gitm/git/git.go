package git

import (
	"github.com/mattcarp12/go-gitm/gitm/config"
	"github.com/mattcarp12/go-gitm/gitm/files"
)

func check() {
	files.AssertInRepo()
	config.AssertNotBare()
}

func Diff() {}

func Remote() {}

func Fetch() {}

func Merge() {}

func Pull() {}

func Push() {}

func Clone() {}

func WriteTree() {}

func UpdateRef() {}
