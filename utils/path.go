package utils

import (
	"io"
	"os"
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
