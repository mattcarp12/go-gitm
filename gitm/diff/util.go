package diff

func (d Diff) keys() []string {
	keys := []string{}
	for k := range d {
		keys = append(keys, k)
	}
	return keys
}