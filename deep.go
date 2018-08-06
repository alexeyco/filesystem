package filesystem

import (
	"os"
	"path/filepath"
)

// SeekerDeep deep file seeker
type SeekerDeep struct {
	*SeekerIn
	root string
}

// SetRoot sets seeker root
func (s *SeekerDeep) SetRoot(root string) {
	s.root = root
}

// Root returns seeker root
func (s *SeekerDeep) Root() string {
	return s.root
}

// Each returns iterator with current seeker
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

// Deep returns deep seeker
func Deep() Seeker {
	return &SeekerDeep{}
}
