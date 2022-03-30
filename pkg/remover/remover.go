package remover

import (
	"os"

	"github.com/wariored/brt/pkg/helpers"
)

func RemoveFile(path string) error {
	if !helpers.FileExists(path) {
		return helpers.ErrorFileDoesNotExist
	}
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func RemoveDir(path string) error {
	if !helpers.DirExists(path) {
		return helpers.ErrorDirDoesNotExist
	}
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}