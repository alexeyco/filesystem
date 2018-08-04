package filesystem

type Dir struct {
	parent *Dir
	name   string
	path   string
	dirs   []*Dir
	files  []*File
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

func (*Dir) Rename(oldName, newName string) error {
	return nil
}

func (*Dir) Move(dir *Dir) error {
	return nil
}

func (*Dir) Remove(name string) error {
	return nil
}

func (d *Dir) List() ([]Path, error) {
	return []Path{}, nil
}

func (d *Dir) Dirs() ([]*Dir, error) {
	return []*Dir{}, nil
}

func (d *Dir) Files() ([]File, error) {
	return []File{}, nil
}

func (d *Dir) Mkdir(name string) (*Dir, error) {
	return &Dir{}, nil
}

func newDir(path string) (*Dir, error) {
	return &Dir{}, nil
}
