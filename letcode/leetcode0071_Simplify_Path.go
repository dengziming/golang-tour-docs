package letcode

import (
	"strings"
)

func simplifyPath(path string) string {
	names := strings.Split(path,"/")

	result := make([]string, 0)
	for _, name := range names {
		if name == "" || name == "." {

		} else if name == ".." {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, name)
		}
	}

	return "/" + strings.Join(result, "/")
}
