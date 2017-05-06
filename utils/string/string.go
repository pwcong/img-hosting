package string

import (
	"strings"
)

func PickFilenameFromContentDisposition(str string) string {

	if len(str) < 1 {
		return ""
	}

	res := strings.Split(str, "filename=\"")

	if len(res) < 2 {
		return ""
	}

	if res[1] == "" {
		return ""
	}

	return res[1][0 : len(res[1])-1]

}
