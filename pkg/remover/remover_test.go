package remover

import (
	"os"
	"testing"

	"github.com/wariored/brt/pkg/adder"
	"github.com/wariored/brt/pkg/helpers"
)

func TestRemoveFile(t *testing.T) {
	path := helpers.GetWorkingDir() + "/tmp/removefile.txt"
	adder.CreateFile(path)
	got := RemoveFile(path)

	if got != nil {
		t.Errorf("RemoveFile(%s), got '%v'; want nil", path, got)
	}
	os.Remove(path)
}

func TestRemoveFile_FileNotExists(t *testing.T) {
	path := helpers.GetWorkingDir() + "/tmp/removefile.txt"
	got := RemoveFile(path)

	if got != helpers.ErrorFileDoesNotExist {
		t.Errorf("RemoveFile(%s), got '%v'; want '%v'", path, got, helpers.ErrorFileDoesNotExist)
	}
}

func TestRemoveDir(t *testing.T) {
	path := helpers.GetWorkingDir() + "/tmp/removedir"
	adder.CreateDir(path)
	got := RemoveDir(path)

	if got != nil {
		t.Errorf("RemoveDir(%s), got '%v'; want nil", path, got)
	}
	os.Remove(path)
}

func TestRemoveDir_DirNotExist(t *testing.T) {
	path := helpers.GetWorkingDir() + "/tmp/removedir"
	got := RemoveDir(path)

	if got != helpers.ErrorDirDoesNotExist {
		t.Errorf("RemoveDir(%s), got '%v'; want '%v'", path, got, helpers.ErrorDirDoesNotExist)
	}
}