package util

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

func Keys(m map[string]string) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Unique(strArr []string) []string {
	strMap := map[string]string{}
	for _, str := range strArr {
		strMap[str] = ""
	}
	return Keys(strMap)
}