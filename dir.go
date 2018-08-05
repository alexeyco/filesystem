package filesystem

import (
	"os"
	"path/filepath"
)

type EachDirHandler func(dir *Dir)

type Dir struct {
	name string
	info os.FileInfo
}

func (d *Dir) IsDir() bool {
	return true
}

func (d *Dir) IsFile() bool {
	return false
}

func (d *Dir) Name() string {
	return d.name
}

func newDir(path string, info os.FileInfo) *Dir {
	return &Dir{
		name: filepath.Join(path, info.Name()),
		info: info,
	}
}
