package models

type Dir struct {
	Name string
	Path string
}

type File struct {
	Dir *Dir
	Name string
}

