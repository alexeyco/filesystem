package filesystem

import "os"

type EachDirHandler func(dir *Dir)

type Dir struct {
	info os.FileInfo
}

func (d *Dir) IsDir() bool {
	return true
}

func (d *Dir) IsFile() bool {
	return false
}

func (*Dir) Name() string {
	return ""
}

func (d *Dir) Stat() os.FileInfo {
	return d.info
}

func newDir(info os.FileInfo) *Dir {
	return &Dir{
		info: info,
	}
}
