package filesystem

import (
	"os"
	"path/filepath"
)

// Fs filesystem object
type Fs struct {
	abs  string // root absolute abs
	root *Dir   // root directory object
}

// Abs returns absolute abs
func (fs *Fs) Abs() string {
	return fs.abs
}

// List returns nested contents
func (fs *Fs) List() (Paths, error) {
	return fs.root.List()
}

// Dirs returns nested directories
func (fs *Fs) Dirs() (Dirs, error) {
	return fs.root.Dirs()
}

// Dir returns nested directory by name
func (fs *Fs) Dir(name string) (*Dir, error) {
	dirs, err := fs.root.Dirs()
	if err != nil {
		return nil, err
	}

	return dirs.Dir(name)
}

// Files returns nested files
func (fs *Fs) Files() (Files, error) {
	return fs.root.Files()
}

// File returns nested file by name
func (fs *Fs) File(name string) (*File, error) {
	files, err := fs.root.Files()
	if err != nil {
		return nil, err
	}

	return files.File(name)
}

// Root returns root directory
func Root(path string) (*Fs, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, &ErrDirNotFound{
			path: path,
		}
	}

	if !stat.IsDir() {
		return nil, &ErrMustBeDirectory{
			path: path,
		}
	}

	fs := &Fs{abs: path}
	fs.root = newDir(fs, nil, "")

	return fs, nil
}
