package filesystem

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// SeekerIn basic seeker
type SeekerIn struct {
	root string
	dir  string
}

// Root returns seeker root
func (s *SeekerIn) Root() string {
	return filepath.Join(s.root, s.dir)
}

// Each returns iterator
func (s *SeekerIn) Each() *Iterator {
	return &Iterator{
		seeker: s,
	}
}

// SetRoot sets seeker root (relative to Fs root)
func (s *SeekerIn) SetRoot(root string) {
	s.root = root
}

// Exist checks if file exist
func (s *SeekerIn) Exist(path string) bool {
	_, err := s.entry(path)
	if err != nil {
		return false
	}

	return true
}

// IsDir checks if path is directory
func (s *SeekerIn) IsDir(path string) bool {
	entry, err := s.entry(path)
	if err != nil {
		return false
	}

	return entry.IsDir()
}

// IsFile checks if path is file
func (s *SeekerIn) IsFile(path string) bool {
	entry, err := s.entry(path)
	if err != nil {
		return false
	}

	return entry.IsFile()
}

// Dir returns directory by path
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

// File returns file by path
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

// Mkdir creates directory
func (s *SeekerIn) Mkdir(path string) error {
	if err := s.shouldNotBeExist(path); err != nil {
		return err
	}

	return os.MkdirAll(filepath.Join(s.root, s.dir, path), os.ModePerm)
}

// Move moves files and directories
func (s *SeekerIn) Move(source, dest string) error {
	if err := s.shouldBeExist(source); err != nil {
		return err
	}

	if err := s.shouldNotBeExist(dest); err != nil {
		return err
	}

	return os.Rename(filepath.Join(s.dir, s.root, source), filepath.Join(s.dir, s.root, dest))
}

// Remove removes files and directories
func (s *SeekerIn) Remove(path string) error {
	if err := s.shouldBeExist(path); err != nil {
		return err
	}

	return os.RemoveAll(filepath.Join(s.root, s.dir, path))
}

func (s *SeekerIn) inRoot(path string) error {
	dir := filepath.Join(s.root, s.dir)
	return inRoot(dir, filepath.Join(dir, path))
}

func (s *SeekerIn) shouldBeExist(path string) error {
	if err := s.inRoot(path); err != nil {
		return err
	}

	if !s.Exist(path) {
		return &ErrNotExist{path: path}
	}

	return nil
}

func (s *SeekerIn) shouldNotBeExist(path string) error {
	if err := s.inRoot(path); err != nil {
		return err
	}

	if s.Exist(path) {
		return &ErrAlreadyExist{path: path}
	}

	return nil
}

func (s *SeekerIn) entry(path string) (Entry, error) {
	if err := s.inRoot(path); err != nil {
		return nil, err
	}

	path = filepath.Join(s.root, s.dir, path)
	local, err := s.stripPath(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err = checkNotExist(local, err); err != nil {
		return nil, err
	}
	defer f.Close()

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
	if checkNotExist(local, err) != nil {
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

// In returns SeekerIn
func In(dir string) Seeker {
	return &SeekerIn{
		dir: dir,
	}
}
