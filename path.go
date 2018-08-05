package filesystem

import (
	"io/ioutil"
)

// Path basic filesystem item
type Path interface {
	IsFile() bool
	IsDir() bool
	Parent() *Dir
	Name() string
	Path() string
	abs() string
	Rename(name string) error
	Move(dir *Dir) error
	Remove() error
}

// HandlerEachPath handler for iteration through all contents
type HandlerEachPath func(path Path)

// Paths directory contents collection
type Paths []Path

// Dirs returns dirs
func (p Paths) Dirs() Dirs {
	dirs := Dirs{}

	for _, path := range p {
		if path.IsDir() {
			d, _ := path.(*Dir)
			dirs[d.Name()] = d
		}
	}

	return dirs
}

// Files returns files
func (p Paths) Files() Files {
	files := Files{}

	for _, path := range p {
		if path.IsFile() {
			f, _ := path.(*File)
			files[f.Name()] = f
		}
	}

	return files
}

// Each iterates through all contents
func (p Paths) Each(handler HandlerEachPath) {
	for _, path := range p {
		handler(path)
	}
}

func newPaths(dir *Dir) (Paths, error) {
	paths, err := ioutil.ReadDir(dir.abs())
	if err != nil {
		return Paths{}, err
	}

	p := Paths{}
	for _, path := range paths {
		if !path.IsDir() {
			p = append(p, newFile(dir, path.Name()))
		} else if path.Name() != "." && path.Name() != ".." {
			p = append(p, newDir(dir.fs, dir, path.Name()))
		}
	}

	return p, nil
}
