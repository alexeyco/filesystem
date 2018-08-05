package filesystem

import (
	"path/filepath"
)

type Fs struct {
	root string
}

func (fs *Fs) Read(seeker Seeker) Seeker {
	seeker.SetRoot(fs.root)
	return seeker
}

func (fs *Fs) Each() *Iterator {
	return fs.Read(In("")).Each()
}

func (fs *Fs) Exists(path string) bool {
	return fs.inRoot().Exists(path)
}

func (fs *Fs) IsDir(path string) bool {
	return fs.inRoot().IsDir(path)
}

func (fs *Fs) IsFile(path string) bool {
	return fs.inRoot().IsFile(path)
}

func (fs *Fs) Dir(path string) (*Dir, error) {
	return fs.inRoot().Dir(path)
}

func (fs *Fs) File(path string) (*File, error) {
	return fs.inRoot().File(path)
}

func (fs *Fs) Mkdir(path string) error {
	return fs.inRoot().Mkdir(path)
}

func (fs *Fs) Rename(source, dest string) error {
	return fs.inRoot().Rename(source, dest)
}

func (fs *Fs) Move(source, dest string) error {
	return fs.inRoot().Remove(source, dest)
}

func (fs *Fs) Remove(path string) error {
	return fs.inRoot().Remove(path)
}

func (fs *Fs) inRoot() Seeker {
	return fs.Read(In(""))
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
