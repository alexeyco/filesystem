package filesystem

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// Dir directory type
type Dir struct {
	fs     *Fs    // root object
	depth  int    // directory depth
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
func (d *Dir) Rename(name string) error {
	d.lock()
	defer d.unlock()

	return d.rename(name)
}

func (d *Dir) rename(name string) error {
	dest := filepath.Join(d.Parent().abs(), name)
	if err := copyAll(d.abs(), dest); err != nil {
		return err
	}

	if err := d.remove(); err != nil {
		return err
	}

	d.name = name
	d.path = filepath.Join(d.Parent().Path(), name)
	d.Parent().flush()

	return nil
}

// Move moves directory with all contents to destination directory
func (d *Dir) Move(dir *Dir) error {
	d.lock()
	defer d.unlock()

	return d.move(dir)
}

func (d *Dir) move(dir *Dir) error {
	dest := filepath.Join(dir.abs(), d.name)
	if err := copyAll(d.abs(), dest); err != nil {
		return err
	}

	if err := d.remove(); err != nil {
		return err
	}

	dir.flush()
	d.Parent().flush()

	d.parent = dir
	d.path = dest

	return nil
}

// Remove removes current directory with all contents
func (d *Dir) Remove() error {
	return d.remove()
}

func (d *Dir) remove() error {
	return os.RemoveAll(d.abs())
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

func (d *Dir) flush() {
	d.lPaths = false
	d.lDirs = false
	d.lFiles = false

	d.paths = Paths{}
	d.dirs = Dirs{}
	d.files = Files{}
}

func (d *Dir) lock() {
	d.muPaths.Lock()
	d.muDirs.Lock()
	d.muFiles.Lock()

	d.flush()
}

func (d *Dir) unlock() {
	d.muPaths.Unlock()
	d.muDirs.Unlock()
	d.muFiles.Unlock()
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

func copyAll(source, dest string) error {
	info, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, info.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {
		s := filepath.Join(source, obj.Name())
		d := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			err = copyAll(s, d)
		} else {
			err = copyFile(s, d)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func copyFile(source, dest string) error {
	s, err := os.Open(source)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	if err == nil {
		info, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, info.Mode())
		}
	}

	return err
}
