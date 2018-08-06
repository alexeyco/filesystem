package filesystem

import (
	"os"
	"path/filepath"
)

// EachDirHandler directories iteration handler
type EachDirHandler func(dir *Dir)

// Dir directory type
type Dir struct {
	name string
	info os.FileInfo
}

// IsDir always true
func (d *Dir) IsDir() bool {
	return true
}

// IsFile always false
func (d *Dir) IsFile() bool {
	return false
}

// Name returns directory name
func (d *Dir) Name() string {
	return d.name
}

func newDir(path string, info os.FileInfo) *Dir {
	return &Dir{
		name: filepath.Join(path, info.Name()),
		info: info,
	}
}
