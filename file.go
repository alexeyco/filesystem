package filesystem

import (
	"os"
)

type EachFileHandler func(file *File)

type File struct {
	info os.FileInfo
}

func (*File) IsDir() bool {
	return false
}

func (*File) IsFile() bool {
	return true
}

func (*File) Name() string {
	return ""
}

func (f *File) Stat() os.FileInfo {
	return f.info
}

func newFile(info os.FileInfo) *File {
	return &File{
		info: info,
	}
}
