package filesystem

import (
	"fmt"
	"path/filepath"
	"sync"
)

type Dir struct {
	fs     *Fs
	parent *Dir
	name   string
	path   string

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

func (d *Dir) IsFile() bool {
	return false
}

func (d *Dir) IsDir() bool {
	return true
}

func (d *Dir) Parent() *Dir {
	return d.parent
}

func (d *Dir) Name() string {
	return d.name
}

func (d *Dir) Path() string {
	return d.path
}

func (d *Dir) fullPath() string {
	return filepath.Join(d.fs.Root(), d.path)
}

func (*Dir) Rename(name string) error {
	return nil
}

func (*Dir) Move(dir *Dir) error {
	return nil
}

func (*Dir) Remove() error {
	return nil
}

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

func (d *Dir) Dir(name string) (*Dir, error) {
	dirs, err := d.Dirs()
	if err != nil {
		return nil, err
	}

	return dirs.Dir(name)
}

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

func (d *Dir) File(name string) (*File, error) {
	files, err := d.Files()
	if err != nil {
		return nil, err
	}

	return files.File(name)
}

func (d *Dir) Mkdir(name string) (*Dir, error) {
	return &Dir{}, nil
}

type ErrMustBeDirectory struct {
	path string
}

func (e ErrMustBeDirectory) Error() string {
	return fmt.Sprintf("path %s must be directory", e.path)
}

type ErrDirNotFound struct {
	path string
}

func (e ErrDirNotFound) Error() string {
	return fmt.Sprintf("directory %s not found", e.path)
}

type HandlerEachDir func(dir *Dir)

type Dirs map[string]*Dir

func (d Dirs) Each(handler HandlerEachDir) {
	for _, dir := range d {
		handler(dir)
	}
}

func (d Dirs) Dir(name string) (*Dir, error) {
	var (
		dir *Dir
		ok  bool
		err error
	)

	if dir, ok = d[name]; !ok {
		err = &ErrDirNotFound{
			path: name,
		}
	}

	return dir, err
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
