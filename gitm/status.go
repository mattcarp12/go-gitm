package gitm

// listing keeps `lines` (prefixed by `heading`) only if it's nonempty
func listing(heading string, lines []string) []string {
	if len(lines) > 0 {
		return append([]string{heading}, lines...)
	} else {
		return []string{}
	}
}

// untracked returns list of files not being tracked by gitm
func untracked() []string {
	return []string{}
}

func Status() {

}
