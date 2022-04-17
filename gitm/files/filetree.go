package files

import (
	"path/filepath"
	"strings"
)

// NestFlatTree takes `tree`, a mapping of file path strings to contents,
// and returns a mapping where each key represents a directory path
// eg {"a/b": "hello"} => {"a": {"b": "hello"}}
func NestFlatTree(flatTree map[string]string) map[string]interface{} {
	nestedTree := map[string]interface{}{}
	for key, value := range flatTree {
		// split key into path components
		path := strings.Split(key, string(filepath.Separator))

		// create nested map
		curr := nestedTree

		// iterate through path components
		for i, component := range path {
			// if last component, set value
			if i == len(path)-1 {
				curr[component] = value
			} else {
				// if not last component, create nested map
				if _, ok := curr[component]; !ok {
					curr[component] = map[string]interface{}{}
				}
				// set current map to nested map
				curr = curr[component].(map[string]interface{})
			}
		}
	}
	return nestedTree
}

func FlattenNestedTree(nestedTree map[string]interface{}) map[string]string {
	flatTree := map[string]string{}
	for key, value := range nestedTree {
		switch v := value.(type) {
		case string:
			flatTree[key] = v
		case map[string]interface{}:
			flatTree = FlattenNestedTree(v)
		}
	}
	return flatTree
} 