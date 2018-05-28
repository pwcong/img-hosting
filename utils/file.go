package utils

import (
	"io"
	"os"
	"strings"
)

func CopyFile(srcPath string, dstPath string) error {

	src, err := os.Open(srcPath)

	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	_, err = io.Copy(dst, src)

	return err

}

func GetExtension(filename string) string {

	t := strings.Split(filename, ".")

	if len(t) < 2 {
		return "unknown"
	} else {
		return t[len(t)-1]
	}

}
