package gitm

type Ref string

func (r Ref) isRef() bool { return false }

func (r Ref) terminalRef() Ref {
	return Ref("HEAD")
}

func (r Ref) hash() {}

func isHeadDetached() bool { return false }

func isCheckedOut(branch string) bool {
	return true
}

func toLocalRef(name string) Ref {
	return Ref("refs/heads/" + name)
}

func toRemoteRef(remote, name string) Ref {
	return Ref("refs/remotes/" + remote + "/" + name)
}

func (r Ref) write(content string) {

}

func (r Ref) rm() {}


