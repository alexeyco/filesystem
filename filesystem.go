package filesystem

import (
	"os"
	"path/filepath"
)

type Fs struct {
	root string
}

func (fs *Fs) Read(source Source) *Collection {
	source.setFs(fs)
	return &Collection{
		source: source,
	}
}

func (fs *Fs) Each() *Iterator {
	return fs.Read(In("")).Each()
}

func (fs *Fs) Exists(path string) bool {
	return false
}

func (fs *Fs) IsFile(path string) bool {
	if !fs.Exists(path) {
		return false
	}

	return false
}

func (fs *Fs) IsDir(path string) bool {
	if !fs.Exists(path) {
		return false
	}

	return false
}

func (fs *Fs) Mkdir(path string) error {
	return os.MkdirAll(fs.abs(path), os.ModePerm)
}

func (fs *Fs) Remove(path string) error {
	return os.RemoveAll(fs.abs(path))
}

func (fs *Fs) Move(source, dest string) error {
	return nil
}

func (fs *Fs) Rename(oldName, newName string) error {
	return nil
}

func (fs *Fs) abs(path string) string {
	return filepath.Join(fs.root, path)
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
