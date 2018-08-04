package filesystem

import (
	"io/ioutil"
	"log"
)

type Path interface {
	IsFile() bool
	IsDir() bool
	Parent() *Dir
	Name() string
	Path() string
	fullPath() string
	Rename(name string) error
	Move(dir *Dir) error
	Remove() error
}

type HandlerEachPath func(path Path)

type Paths []Path

func (p Paths) Dirs() Dirs {
	dirs := Dirs{}

	for _, path := range p {
		if path.IsDir() {
			d, _ := path.(*Dir)
			dirs = append(dirs, d)
		}
	}

	return dirs
}

func (p Paths) Files() Files {
	files := Files{}

	for _, path := range p {
		if path.IsFile() {
			d, _ := path.(*File)
			files = append(files, d)
		}
	}

	return files
}

func (p Paths) Each(handler HandlerEachPath) {
	for _, path := range p {
		handler(path)
	}
}

func newPaths(dir *Dir) (Paths, error) {
	paths, err := ioutil.ReadDir(dir.fullPath())
	if err != nil {
		log.Fatal(err)
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
