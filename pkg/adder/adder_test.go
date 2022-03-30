package adder

import (
	"os"
	"testing"

	"github.com/wariored/brt/pkg/helpers"
)

func TestCreateFile(t *testing.T) {
	path := helpers.GetWorkingDir() + "/tmp/createfile.txt"
	got := CreateFile(path)

	if got != nil {
		t.Errorf("CreateFile(path), got '%v'; want nil", got)
	}
	os.Remove(path)
}

func TestCreateFileAlreadyExists(t *testing.T) {
	path := helpers.GetWorkingDir() + "/tmp/createfile.txt"
	CreateFile(path)
	got := CreateFile(path)

	if got != helpers.ErrorFileAlreadyExists {
		t.Errorf("CreateFile(path), got '%v'; want '%v'", got, helpers.ErrorFileAlreadyExists)
	}
	os.Remove(path)
}

func TestCreateDir(t *testing.T) {
	path := helpers.GetWorkingDir() + "/tmp/createdir"
	got := CreateDir(path)

	if got != nil {
		t.Errorf("CreateDir(path), got '%v'; want nil", got)
	}
	os.Remove(path)
}

func TestCreateDirAlreadyExists(t *testing.T) {
	path := helpers.GetWorkingDir() + "/tmp/createdir"
	CreateDir(path)
	got := CreateDir(path)

	if got != helpers.ErrorDirAlreadyExists {
		t.Errorf("CreateDir(path), got '%v'; want '%v'", got, helpers.ErrorDirAlreadyExists)
	}
	os.Remove(path)
}
