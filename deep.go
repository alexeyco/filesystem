package filesystem

import (
	"os"
	"path/filepath"
)

type SeekerDeep struct {
	*SeekerIn
	root string
}

func (s *SeekerDeep) SetRoot(root string) {
	s.root = root
}

func (s *SeekerDeep) Root() string {
	return s.root
}

func (s *SeekerDeep) Each() *Iterator {
	return &Iterator{
		seeker: s,
	}
}

func (s *SeekerDeep) each(handler EachEntryHandler) error {
	return filepath.Walk(s.root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if err := inRoot(s.root, path); err != nil {
			return nil
		}

		local, err := filepath.Rel(s.root, filepath.Dir(path))
		if err != nil {
			return err
		}

		handler(newEntry(local, info))
		return err
	})
}

func Deep() Seeker {
	return &SeekerDeep{}
}
