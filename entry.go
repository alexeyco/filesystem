package filesystem

import (
	"fmt"
	"os"
)

type EachEntryHandler func(entry Entry)

type Entry interface {
	IsDir() bool
	IsFile() bool
	Name() string
}

func newEntry(name string, info os.FileInfo) Entry {
	if info.IsDir() {
		return newDir(name, info)
	} else {
		return newFile(name, info)
	}
}

type ErrNotExists struct {
	path string
}

func (e *ErrNotExists) Error() string {
	return fmt.Sprintf("in %s not exists", e.path)
}
