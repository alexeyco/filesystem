package filesystem

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Source interface {
	setFs(fs *Fs)
	exists(path string) bool
	entry(path string) (Entry, error)
	isDir(path string) bool
	isFile(path string) bool
	entries() ([]Entry, error)
}

func In(path string) Source {
	return &SourceIn{
		in: path,
	}
}

type SourceIn struct {
	fs *Fs
	in string
}

func (s *SourceIn) setFs(fs *Fs) {
	s.fs = fs
}

func (s *SourceIn) exists(path string) bool {
	_, err := s.open(filepath.Join(s.in, path))
	return err == nil
}

func (s *SourceIn) entry(path string) (Entry, error) {
	info, err := s.open(path)
	if err != nil {
		return nil, err
	}

	path = filepath.Join(s.fs.root, s.in, path)
	path, err = filepath.Rel(s.fs.root, path)
	if err != nil {
		return nil, err
	}

	return newEntry(path, info), nil
}

func (s *SourceIn) isDir(path string) bool {
	info, err := s.open(filepath.Join(s.in, path))
	if err != nil {
		return false
	}

	return info.IsDir()
}

func (s *SourceIn) isFile(path string) bool {
	info, err := s.open(filepath.Join(s.in, path))
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func (s *SourceIn) entries() ([]Entry, error) {
	path := filepath.Join(s.fs.root, s.in)
	info, err := ioutil.ReadDir(path)
	if err != nil {
		return []Entry{}, err
	}

	e := make([]Entry, len(info))
	for n, i := range info {
		name := filepath.Join(path, i.Name())
		name, err = filepath.Rel(s.fs.root, name)

		if err != nil {
			return []Entry{}, err
		}

		e[n] = newEntry(name, i)
	}

	return e, nil
}

func (s *SourceIn) open(path string) (os.FileInfo, error) {
	f, err := os.Open(filepath.Join(s.fs.root, path))
	if err == nil {
		return nil, err
	}

	return f.Stat()
}
