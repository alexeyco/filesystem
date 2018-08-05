package filesystem

import (
	"fmt"
	"os"
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
func (f *File) Rename(name string) error {
	p := f.Parent()

	p.lock()
	defer p.unlock()

	oldName := f.abs()
	newName := filepath.Join(p.abs(), name)

	err := os.Rename(oldName, newName)
	if err != nil {
		f.name = name
		f.path = filepath.Join(f.Parent().Path(), name)
	}

	return err
}

// Move moves current file to destination directory
func (f *File) Move(dir *Dir) error {
	p := f.Parent()

	p.lock()
	dir.lock()

	defer p.unlock()
	defer dir.unlock()

	return nil
}

// Remove removes current file
func (f *File) Remove() error {
	p := f.Parent()

	p.lock()
	defer p.unlock()

	return os.RemoveAll(f.abs())
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

// Exists checks if nested file exists
func (f Files) Exists(name string) bool {
	_, ok := f[name]
	return ok
}

// File returns file by name
func (f Files) File(name string) (*File, error) {
	if !f.Exists(name) {
		return nil, &ErrFileNotFound{
			path: name,
		}
	}

	return f[name], nil
}

func newFile(parent *Dir, local string) *File {
	return &File{
		parent: parent,
		name:   local,
		path:   filepath.Join(parent.path, local),
	}
}
