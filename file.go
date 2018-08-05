package filesystem

import (
	"fmt"
	"path/filepath"
)

// File file object
type File struct {
	parent *Dir   // parent directory
	name   string // file name
	path   string // path from root
}

// IsFile always true
func (*File) IsFile() bool {
	return true
}

// IsDir always false
func (*File) IsDir() bool {
	return false
}

// Parent returns parent directory
func (f *File) Parent() *Dir {
	return f.parent
}

// Name returns file name
func (f *File) Name() string {
	return f.name
}

// Path returns file abs
func (f *File) Path() string {
	return f.path
}

// abs returns file absolute abs
func (f *File) abs() string {
	return filepath.Join(f.Parent().abs(), f.Name())
}

// Rename renames current file
func (*File) Rename(name string) error {
	return nil
}

// Move moves current file to destination directory
func (*File) Move(dir *Dir) error {
	return nil
}

// Remove removes current file
func (*File) Remove() error {
	return nil
}

// ErrFileNotFound file not found
type ErrFileNotFound struct {
	path string
}

// Error returns error string
func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("file %s not found", e.path)
}

// HandlerEachFile handler for iteration through files
type HandlerEachFile func(file *File)

// Files files collection
type Files map[string]*File

// Each iterates through files
func (f Files) Each(handler HandlerEachFile) {
	for _, file := range f {
		handler(file)
	}
}

// File returns file by name
func (f Files) File(name string) (*File, error) {
	var (
		file *File
		ok   bool
		err  error
	)

	if file, ok = f[name]; !ok {
		err = &ErrFileNotFound{
			path: name,
		}
	}

	return file, err

}

func newFile(parent *Dir, local string) *File {
	return &File{
		parent: parent,
		name:   local,
		path:   filepath.Join(parent.path, local),
	}
}