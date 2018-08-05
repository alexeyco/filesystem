package filesystem

type Seeker interface {
	Each() *Iterator
	Exists(path string) bool
	IsDir(path string) bool
	IsFile(path string) bool
	Dir(path string) (*Dir, error)
	File(path string) (*File, error)
	Mkdir(path string) error
	Rename(source, dest string) error
	Move(source, dest string) error
	Remove(path string) error
	SetRoot(string)
}

type SeekerIn struct {
	root string
	dir  string
}

func (s *SeekerIn) Each() *Iterator {
	return &Iterator{
		seeker: s,
	}
}

func (s *SeekerIn) Exists(path string) bool {
	panic("implement me")
}

func (s *SeekerIn) IsDir(path string) bool {
	panic("implement me")
}

func (s *SeekerIn) IsFile(path string) bool {
	panic("implement me")
}

func (s *SeekerIn) Dir(path string) (*Dir, error) {
	panic("implement me")
}

func (s *SeekerIn) File(path string) (*File, error) {
	panic("implement me")
}

func (s *SeekerIn) Mkdir(path string) error {
	panic("implement me")
}

func (s *SeekerIn) Rename(source, dest string) error {
	panic("implement me")
}

func (s *SeekerIn) Move(source, dest string) error {
	panic("implement me")
}

func (s *SeekerIn) Remove(path string) error {
	panic("implement me")
}

func (s *SeekerIn) SetRoot(root string) {
	s.root = root
}

func In(dir string) Seeker {
	return &SeekerIn{
		dir: dir,
	}
}
