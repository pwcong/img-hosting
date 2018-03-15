package utils

import (
	"os"
)

func MkdirIFNotExist(dirPath string) error {

	_, err := os.Stat(dirPath)

	if err == nil {
		return nil
	}

	err = os.Mkdir(dirPath, 0666)
	if err != nil {
		return err
	}

	return nil
}
