package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

func GetWorkingDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

func GetFileExt(basename string) string {
	return filepath.Ext(basename)
}

func GetFileCompletePath(path string) string {
	if strings.HasPrefix(path, "/"){
		return path
	}
	return filepath.Join(GetWorkingDir(), path)

}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func FileHasExtension(path string) bool {
	return strings.Contains(path, ".")
}