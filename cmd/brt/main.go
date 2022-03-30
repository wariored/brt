package main

import (
	"os"

	"github.com/gookit/color"
	"github.com/teris-io/cli"
	"github.com/wariored/brt/pkg/adder"
	"github.com/wariored/brt/pkg/helpers"
	"github.com/wariored/brt/pkg/remover"
	"github.com/wariored/brt/pkg/substitor"
)

func main() {
	addFile := cli.NewCommand("add-file", "Add a file").
		WithShortcut("af").
		WithArg(cli.NewArg("filename", "filename to add")).
		WithAction(func(args []string, options map[string]string) int {
			path := helpers.GetFileCompletePath(args[0])
			err := adder.CreateFile(path)
			if err != nil {
				color.Red.Println(err)
				return 0
			}
			color.Green.Printf("file added, location: %s", path)
			return 0
		})
	addDir := cli.NewCommand("add-dir", "Add a directory ").
		WithShortcut("ad").
		WithArg(cli.NewArg("dirname", "dirname to add")).
		WithAction(func(args []string, options map[string]string) int {
			path := helpers.GetFileCompletePath(args[0])
			err := adder.CreateDir(path)
			if err != nil {
				color.Red.Println(err)
				return 0
			}
			color.Green.Printf("directory added, location: %s", path)
			return 0
		})
	copyFile := cli.NewCommand("copy-file", "Copy a file").
		WithShortcut("cf").
		WithArg(cli.NewArg("source", "filename to copy")).
		WithArg(cli.NewArg("destination", "filename to copy to")).
		WithAction(func(args []string, options map[string]string) int {
			src := helpers.GetFileCompletePath(args[0])
			dst := helpers.GetFileCompletePath(args[1])
			err := substitor.CopyFile(src, dst)
			if err != nil {
				color.Red.Println(err)
			}
			return 0
		})
	copyDir := cli.NewCommand("copy-dir", "Copy a directory").
		WithShortcut("cd").
		WithArg(cli.NewArg("source", "dirname to copy")).
		WithArg(cli.NewArg("destination", "dirname to copy to")).
		WithAction(func(args []string, options map[string]string) int {
			src := helpers.GetFileCompletePath(args[0])
			dst := helpers.GetFileCompletePath(args[1])
			err := substitor.CopyDir(src, dst)
			if err != nil {
				color.Red.Println(err)
			}
			return 0
		})
	removeFile := cli.NewCommand("remove-file", "Remove a file").
		WithShortcut("rf").
		WithArg(cli.NewArg("filename", "filename to remove")).
		WithAction(func(args []string, options map[string]string) int {
			path := helpers.GetFileCompletePath(args[0])
			err := remover.RemoveFile(path)
			if err != nil {
				color.Red.Println(err)
				return 0
			}
			color.Green.Printf("file %s removed", path)
			return 0
		})
	removeDir := cli.NewCommand("remove-dir", "Remove a directory").
		WithShortcut("rd").
		WithArg(cli.NewArg("dirname", "dirname to delete")).
		WithAction(func(args []string, options map[string]string) int {
			path := helpers.GetFileCompletePath(args[0])
			err := remover.RemoveDir(path)
			if err != nil {
				color.Red.Println(err)
				return 0
			}
			color.Green.Printf("directory %s removed", path)
			return 0
	})
	app := cli.New("File management tool").
		WithCommand(addFile).
		WithCommand(addDir).
		WithCommand(copyFile).
		WithCommand(copyDir).
		WithCommand(removeDir).
		WithCommand(removeFile)

	os.Exit(app.Run(os.Args, os.Stdout))
}
