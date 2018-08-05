package filesystem

import (
	"fmt"
	"path/filepath"
	"sync"
)

// Dir directory type
type Dir struct {
	fs     *Fs    // root object
	parent *Dir   // parent directory
	name   string // directory name
	path   string // path from root directory

	lPaths  bool
	muPaths sync.Mutex
	paths   Paths

	lDirs  bool
	muDirs sync.Mutex
	dirs   Dirs

	lFiles  bool
	muFiles sync.Mutex
	files   Files
}

// IsFile always false
func (d *Dir) IsFile() bool {
	return false
}

// IsDir always true
func (d *Dir) IsDir() bool {
	return true
}

// Parent returns parent directory object; if directory is root - returns nil
func (d *Dir) Parent() *Dir {
	return d.parent
}

// Name returns directory name
func (d *Dir) Name() string {
	return d.name
}

// Path returns directory abs
func (d *Dir) Path() string {
	return d.path
}

// abs returns directory absolute abs
func (d *Dir) abs() string {
	return filepath.Join(d.fs.Abs(), d.path)
}

// Rename renames current directory
func (*Dir) Rename(name string) error {
	return nil
}

// Move moves directory with all contents to destination directory
func (*Dir) Move(dir *Dir) error {
	return nil
}

// Remove removes current directory with all contents
func (*Dir) Remove() error {
	return nil
}

// List returns all directory nested contents
func (d *Dir) List() (Paths, error) {
	d.muPaths.Lock()
	defer d.muPaths.Unlock()

	if d.lPaths {
		return d.paths, nil
	}

	p, err := newPaths(d)
	if err != nil {
		return Paths{}, err
	}

	d.lPaths = true
	d.paths = p

	return d.paths, nil
}

// Dirs returns nested directories list
func (d *Dir) Dirs() (Dirs, error) {
	d.muDirs.Lock()
	defer d.muDirs.Unlock()

	if d.lDirs {
		return d.dirs, nil
	}

	p, err := d.List()
	if err != nil {
		return Dirs{}, err
	}

	d.lDirs = true
	d.dirs = p.Dirs()

	return d.dirs, nil
}

// Dir returns nested directory by name
func (d *Dir) Dir(name string) (*Dir, error) {
	dirs, err := d.Dirs()
	if err != nil {
		return nil, err
	}

	return dirs.Dir(name)
}

// Files returns nested files
func (d *Dir) Files() (Files, error) {
	d.muFiles.Lock()
	defer d.muFiles.Unlock()

	if d.lFiles {
		return d.files, nil
	}

	p, err := d.List()
	if err != nil {
		return Files{}, err
	}

	d.lFiles = true
	d.files = p.Files()

	return d.files, nil
}

// File returns nested file by name
func (d *Dir) File(name string) (*File, error) {
	files, err := d.Files()
	if err != nil {
		return nil, err
	}

	return files.File(name)
}

// Mkdir creates a new nested directory
func (d *Dir) Mkdir(name string) (*Dir, error) {
	return &Dir{}, nil
}

// ErrMustBeDirectory error must be directory
type ErrMustBeDirectory struct {
	path string
}

// Error returns error string
func (e ErrMustBeDirectory) Error() string {
	return fmt.Sprintf("abs %s must be directory", e.path)
}

// ErrDirNotFound error directory not found
type ErrDirNotFound struct {
	path string
}

// Error returns error string
func (e ErrDirNotFound) Error() string {
	return fmt.Sprintf("directory %s not found", e.path)
}

// HandlerEachDir handler for iteration through directories
type HandlerEachDir func(dir *Dir)

// Dirs directories collection
type Dirs map[string]*Dir

// Each iterates through directories
func (d Dirs) Each(handler HandlerEachDir) {
	for _, dir := range d {
		handler(dir)
	}
}

// Exists checks if nested directory exists
func (d Dirs) Exists(name string) bool {
	_, ok := d[name]
	return ok
}

// Dir returns nested directory by name
func (d Dirs) Dir(name string) (*Dir, error) {
	if !d.Exists(name) {
		return nil, &ErrDirNotFound{
			path: name,
		}
	}

	return d[name], nil
}

func newDir(fs *Fs, parent *Dir, local string) *Dir {
	path := local
	if parent != nil {
		path = filepath.Join(parent.path, path)
	}

	return &Dir{
		fs:     fs,
		parent: parent,
		name:   local,
		path:   path,
	}
}
