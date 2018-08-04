package filesystem

type File struct {
	parent *Dir
	name   string
	path   string
}

func (*File) IsFile() bool {
	return true
}

func (*File) IsDir() bool {
	return false
}

func (f *File) Parent() *Dir {
	return f.parent
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Path() string {
	return f.path
}

func (*File) Rename(oldName, newName string) error {
	return nil
}

func (*File) Move(dir *Dir) error {
	return nil
}

func (*File) Remove(name string) error {
	return nil
}
