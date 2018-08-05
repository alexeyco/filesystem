package filesystem

import (
	"os"
	"path/filepath"
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

func newFile(path string, info os.FileInfo) *File {
	return &File{
		name: filepath.Join(path, info.Name()),
		info: info,
	}
}
