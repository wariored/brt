package substitor

import (
	"io/ioutil"
	"os"

	"github.com/wariored/brt/pkg/helpers"
)

func CopyFile(src string, dst string) error {
	if !helpers.FileExists(src) {
		return helpers.ErrorFileDoesNotExist
	}
	if helpers.FileExists(dst) {
		return helpers.ErrorFileAlreadyExists
	}
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, content, 0755)
	if err != nil {
		return err
	}
	return nil
}

func CopyDir(src string, dst string) error {
	if !helpers.DirExists(src) {
		return helpers.ErrorDirDoesNotExist
	}
	if !helpers.DirExists(dst) {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}
	

	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			err = CopyDir(src+"/"+f.Name(), dst+"/"+f.Name())
			if err != nil {
				return err
			}
		} else {
			err := CopyFile(src+"/"+f.Name(), dst+"/"+f.Name())
			if err != nil {
				return err

			}

		}
	}
	return nil
}
