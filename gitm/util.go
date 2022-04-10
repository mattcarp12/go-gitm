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

func StringIndex(strArr []string, find string) int {
	for i, str := range strArr {
		if find == str {
			return i
		}
	}
	return -1
}