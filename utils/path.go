package utils

import (
	"os"
	"path/filepath"
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

func GetRootDir() string {

	rootDir, err := os.Getwd()
	if err != nil {
		rootDir = filepath.Dir(os.Args[0])
	}

	return rootDir
}
