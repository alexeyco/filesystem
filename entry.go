package filesystem

import "os"

type EachEntryHandler func(entry Entry)

type Entry interface {
	IsDir() bool
	IsFile() bool
	Name() string
	Stat() os.FileInfo
}

func newEntry(info os.FileInfo) Entry {
	if info.IsDir() {
		return newDir(info)
	} else {
		return newFile(info)
	}
}
