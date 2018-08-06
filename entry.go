package filesystem

import (
	"os"
)

// EachEntryHandler files iteration handler
type EachEntryHandler func(entry Entry)

// Entry basic filesystem object
type Entry interface {
	IsDir() bool
	IsFile() bool
	Name() string
}

func newEntry(name string, info os.FileInfo) Entry {
	if info.IsDir() {
		return newDir(name, info)
	}

	return newFile(name, info)
}
