package gitm

func Intersection(a, b []string) []string {
	res := []string{}

	for _, vala := range a {
		for _, valb := range b {
			if vala == valb {
				res = append(res, vala)
			}
		}
	}

	return res
}
