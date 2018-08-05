package filesystem

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Seeker interface {
	SetRoot(string)
	Root() string
	Each() *Iterator
	Exist(path string) bool
	IsDir(path string) bool
	IsFile(path string) bool
	Dir(path string) (*Dir, error)
	File(path string) (*File, error)
	Mkdir(path string) error
	Move(source, dest string) error
	Remove(path string) error
	each(handler EachEntryHandler) error
}

type SeekerIn struct {
	root string
	dir  string
}

func (s *SeekerIn) Root() string {
	return filepath.Join(s.root, s.dir)
}

func (s *SeekerIn) Each() *Iterator {
	return &Iterator{
		seeker: s,
	}
}

func (s *SeekerIn) SetRoot(root string) {
	s.root = root
}

func (s *SeekerIn) Exist(path string) bool {
	_, err := s.entry(path)
	if err != nil {
		return false
	}

	return true
}

func (s *SeekerIn) IsDir(path string) bool {
	entry, err := s.entry(path)
	if err != nil {
		return false
	}

	return entry.IsDir()
}

func (s *SeekerIn) IsFile(path string) bool {
	entry, err := s.entry(path)
	if err != nil {
		return false
	}

	return entry.IsFile()
}

func (s *SeekerIn) Dir(path string) (*Dir, error) {
	entry, err := s.entry(path)
	if err != nil {
		return nil, err
	}

	if !entry.IsDir() {
		return nil, &ErrNotDir{path: entry.Name()}
	}

	return entry.(*Dir), nil
}

func (s *SeekerIn) File(path string) (*File, error) {
	entry, err := s.entry(path)
	if err != nil {
		return nil, err
	}

	if !entry.IsFile() {
		return nil, &ErrNotFile{path: entry.Name()}
	}

	return entry.(*File), nil
}

func (s *SeekerIn) Mkdir(path string) error {
	path = filepath.Join(s.root, s.dir, path)
	return os.MkdirAll(path, os.ModePerm)
}

func (s *SeekerIn) Move(source, dest string) error {
	source = filepath.Join(s.dir, s.root, source)
	dest = filepath.Join(s.dir, s.root, dest)

	return os.Rename(source, dest)
}

func (s *SeekerIn) Remove(path string) error {
	path = filepath.Join(s.root, s.dir, path)
	return os.RemoveAll(path)
}

func (s *SeekerIn) entry(path string) (Entry, error) {
	path = filepath.Join(s.root, s.dir, path)
	local, err := s.stripPath(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if checkNotNotExist(local, err) != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	entry := newEntry(local, stat)
	return entry, nil
}

func (s *SeekerIn) each(handler EachEntryHandler) error {
	path := filepath.Join(s.root, s.dir)
	local, err := s.stripPath(path)
	if err != nil {
		return err
	}

	info, err := ioutil.ReadDir(path)
	if checkNotNotExist(local, err) != nil {
		return err
	}

	for _, i := range info {
		handler(newEntry(local, i))
	}

	return nil
}

func (s *SeekerIn) stripPath(path string) (string, error) {
	return filepath.Rel(s.root, path)
}

func In(dir string) Seeker {
	return &SeekerIn{
		dir: dir,
	}
}

func checkNotNotExist(path string, err error) error {
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		return &ErrNotExists{
			path: path,
		}
	}

	return err
}
