package filesystem

import (
	"path/filepath"
)

type File struct {
	parent *Dir
	name   string
	path   string
}

func (*File) IsFile() bool {
	return true
}

func (*File) IsDir() bool {
	return false
}

func (f *File) Parent() *Dir {
	return f.parent
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Path() string {
	return f.path
}

func (f *File) fullPath() string {
	return filepath.Join(f.Parent().fullPath(), f.Name())
}

func (*File) Rename(name string) error {
	return nil
}

func (*File) Move(dir *Dir) error {
	return nil
}

func (*File) Remove() error {
	return nil
}

type HandlerEachFile func(file *File)

type Files []*File

func (f Files) Each(handler HandlerEachFile) {
	for _, file := range f {
		handler(file)
	}
}

func newFile(parent *Dir, local string) *File {
	return &File{
		parent: parent,
		name:   local,
		path:   filepath.Join(parent.path, local),
	}
}
