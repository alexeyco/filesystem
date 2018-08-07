package filesystem

import (
	"os"
	"path/filepath"
)

// EachFileHandler file iteration handler
type EachFileHandler func(file *File)

// File file type
type File struct {
	name string
	info os.FileInfo
}

// IsDir always false
func (*File) IsDir() bool {
	return false
}

// IsFile always true
func (*File) IsFile() bool {
	return true
}

// Name returns file name
func (f *File) Name() string {
	return f.name
}

// Ext returns file extension
func (f *File) Ext() string {
	return filepath.Ext(f.Name())
}

// Stat returns file info
func (f *File) Stat() os.FileInfo {
	return f.info
}

func newFile(path string, info os.FileInfo) *File {
	return &File{
		name: filepath.Join(path, info.Name()),
		info: info,
	}
}
