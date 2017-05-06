package file

import "strings"

func GetExtension(filename string) string {

	if len(filename) < 1 {
		return ""
	}

	res := strings.Split(filename, ".")

	if len(res) < 2 {
		return ""
	}

	return res[1]
}
