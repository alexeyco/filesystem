package filesystem

import (
	"io/ioutil"
	"path/filepath"
)

type Source interface {
	setFs(fs *Fs)
	entries() ([]Entry, error)
}

func In(dir string) Source {
	return &SourceIn{
		dir: dir,
	}
}

type SourceIn struct {
	fs  *Fs
	dir string
}

func (s *SourceIn) setFs(fs *Fs) {
	s.fs = fs
}

func (s *SourceIn) entries() ([]Entry, error) {
	root := filepath.Join(s.fs.root, s.dir)
	info, err := ioutil.ReadDir(root)
	if err != nil {
		return []Entry{}, err
	}

	e := make([]Entry, len(info))
	for n, i := range info {
		name := filepath.Join(root, i.Name())
		name, err = filepath.Rel(s.fs.root, name)

		e[n] = newEntry(name, i)
	}

	return e, nil
}
