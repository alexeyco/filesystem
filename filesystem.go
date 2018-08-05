package filesystem

import "path/filepath"

type Fs struct {
	root string
}

func (fs *Fs) Each() *Iterator {
	return fs.In("").Each()
}

func (fs *Fs) In(dir string) Chain {
	return &InChain{
		fs:  fs,
		dir: dir,
	}
}

func (fs *Fs) Mkdir(path string) error {
	return nil
}

func (fs *Fs) Remove(path string) error {
	return nil
}

func (fs *Fs) Move(source, dest string) error {
	return nil
}

func (fs *Fs) Rename(oldName, newName string) error {
	return nil
}

func (fs *Fs) isNested(path string) bool {
	return false
}

func Root(dir string) (*Fs, error) {
	root, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	return &Fs{
		root: root,
	}, nil
}
