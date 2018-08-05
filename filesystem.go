package filesystem

import (
	"os"
	"path/filepath"
)

// Fs root directory object
type Fs struct {
	root string
}

// Root returns absolute path to current root directory
func (fs *Fs) Root() string {
	return fs.root
}

// Read seekers factory
func (fs *Fs) Read(seeker Seeker) Seeker {
	seeker.SetRoot(fs.root)
	return seeker
}

// Each creates an iterator and returns it
func (fs *Fs) Each() *Iterator {
	return fs.Read(In("")).Each()
}

// Exist checks if root-relative path exist
func (fs *Fs) Exist(path string) bool {
	return fs.inRoot().Exist(path)
}

// IsDir checks if root-relative path is a directory
func (fs *Fs) IsDir(path string) bool {
	return fs.inRoot().IsDir(path)
}

// IsFile checks if root-relative path is a file
func (fs *Fs) IsFile(path string) bool {
	return fs.inRoot().IsFile(path)
}

// Dir returns directory info by root-relative path
func (fs *Fs) Dir(path string) (*Dir, error) {
	return fs.inRoot().Dir(path)
}

// File returns file info by root-relative path
func (fs *Fs) File(path string) (*File, error) {
	return fs.inRoot().File(path)
}

// Mkdir creates a new directory in root directory
func (fs *Fs) Mkdir(path string) error {
	return fs.inRoot().Mkdir(path)
}

// Move moves anything inside root
func (fs *Fs) Move(source, dest string) error {
	return fs.inRoot().Move(source, dest)
}

// Remove removes root-relative directory or file by path
func (fs *Fs) Remove(path string) error {
	return fs.inRoot().Remove(path)
}

func (fs *Fs) inRoot() Seeker {
	return fs.Read(In(""))
}

// Root returns Fs object by path
func Root(dir string) (*Fs, error) {
	root, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(root)
	if err = checkNotNotExist(dir, err); err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if !stat.IsDir() {
		return nil, &ErrNotDir{path: dir}
	}

	return &Fs{
		root: root,
	}, nil
}
