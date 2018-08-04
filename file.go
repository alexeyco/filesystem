package filesystem

import (
	"fmt"
	"path/filepath"
)

const (
	FileTypeDefault = iota
	FileTypeEditable
	FileTypeImage
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

func (*File) Type() int {
	return FileTypeDefault
}

type ErrFileNotFound struct {
	path string
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("file %s not found", e.path)
}

type HandlerEachFile func(file *File)

type Files map[string]*File

func (f Files) Each(handler HandlerEachFile) {
	for _, file := range f {
		handler(file)
	}
}

func (f Files) File(name string) (*File, error) {
	var (
		file *File
		ok   bool
		err  error
	)

	if file, ok = f[name]; !ok {
		err = &ErrFileNotFound{
			path: name,
		}
	}

	return file, err

}

func newFile(parent *Dir, local string) *File {
	return &File{
		parent: parent,
		name:   local,
		path:   filepath.Join(parent.path, local),
	}
}
