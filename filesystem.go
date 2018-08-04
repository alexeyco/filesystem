package filesystem

type Fs struct {
	root *Dir
}

func (fs *Fs) List() ([]Path, error) {
	return fs.root.List()
}

func (fs *Fs) Dirs() ([]*Dir, error) {
	return fs.root.Dirs()
}

func (fs *Fs) Files() ([]File, error) {
	return fs.root.Files()
}

func Root(path string) (*Fs, error) {
	root, err := newDir(path)
	if err != nil {
		return nil, err
	}

	return &Fs{
		root: root,
	}, nil
}
