package helpers

import "errors"

var (
	ErrorFileAlreadyExists error = errors.New("the file already exists")
	ErrorFileDoesNotExist error = errors.New("the file does not exists")
	ErrorFileDoesntHaveExtension error = errors.New("the file name doesn't have any extension")
	ErrorDirAlreadyExists error = errors.New("the directory already exists")
	ErrorDirDoesNotExist error = errors.New("the directory doesn't exists")
)
