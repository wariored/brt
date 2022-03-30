package substitor

import (
	"fmt"
	"os"
	"testing"

	"github.com/wariored/brt/pkg/adder"
	"github.com/wariored/brt/pkg/helpers"
)


func TestCopyDir(t *testing.T) {
	path := helpers.GetWorkingDir() + "/tmp/copydir"
	targePath := helpers.GetWorkingDir() + "/tmp/copydir_target"
	err := adder.CreateDir(path)
	if err != nil {
		panic(err)
	}
	err  = adder.CreateFile(path + "/file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	got := CopyDir(path, targePath)
	if got != nil {
		t.Errorf("CopyDir(%s, %s+'copydir'), got '%v'; want nil", path, path, got)
	}
	os.RemoveAll(path)
	os.RemoveAll(targePath)
}