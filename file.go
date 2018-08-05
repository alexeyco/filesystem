package filesystem

import (
	"fmt"
	"os"
)

type EachFileHandler func(file *File)

type File struct {
	name string
	info os.FileInfo
}

func (*File) IsDir() bool {
	return false
}

func (*File) IsFile() bool {
	return true
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Stat() os.FileInfo {
	return f.info
}

func newFile(name string, info os.FileInfo) *File {
	return &File{
		name: name,
		info: info,
	}
}

type ErrNotFile struct {
	path string
}

func (e *ErrNotFile) Error() string {
	return fmt.Sprintf("path %s is not a file", e.path)
}
