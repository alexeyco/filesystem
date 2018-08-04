package filesystem

import "path/filepath"

type Fs struct {
	path string
	root *Dir
}

func (fs *Fs) Root() string {
	return fs.path
}

func (fs *Fs) List() (Paths, error) {
	return fs.root.List()
}

func (fs *Fs) Dirs() (Dirs, error) {
	return fs.root.Dirs()
}

func (fs *Fs) Files() (Files, error) {
	return fs.root.Files()
}

func Root(path string) (*Fs, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// TODO: check path is directory

	fs := &Fs{path: abs}
	fs.root = newDir(fs, nil, "")

	return fs, nil
}
