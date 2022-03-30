package adder

import (
	"os"

	"github.com/wariored/brt/pkg/helpers"
)

func CreateFile(path string) error {
	if helpers.FileExists(path) {
		return helpers.ErrorFileAlreadyExists
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func CreateDir(path string) error {
	if helpers.DirExists(path) {
		return helpers.ErrorDirAlreadyExists
	}
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}
