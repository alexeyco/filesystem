package filesystem

import (
	"io/ioutil"
	"path/filepath"
)

type Chain interface {
	All() ([]Entry, error)
	Dirs() ([]*Dir, error)
	Files() ([]*File, error)
	Each() *Iterator
}

type InChain struct {
	fs  *Fs
	dir string
}

func (c *InChain) All() ([]Entry, error) {
	info, err := ioutil.ReadDir(filepath.Join(c.fs.root, c.dir))
	if err != nil {
		return []Entry{}, err
	}

	e := []Entry{}
	for _, i := range info {
		e = append(e, newEntry(i))
	}

	return e, nil
}

func (c *InChain) Dirs() ([]*Dir, error) {
	dirs := []*Dir{}
	c.Each().Entry(func(entry Entry) {
		if entry.IsDir() {
			e, _ := entry.(*Dir)
			dirs = append(dirs, e)
		}
	})

	return dirs, nil
}

func (c *InChain) Files() ([]*File, error) {
	files := []*File{}
	c.Each().Entry(func(entry Entry) {
		if entry.IsFile() {
			f, _ := entry.(*File)
			files = append(files, f)
		}
	})

	return files, nil
}

func (c *InChain) Each() *Iterator {
	return &Iterator{
		chain: c,
	}
}

type Iterator struct {
	chain Chain
}

func (i *Iterator) Entry(handler EachEntryHandler) error {
	e, err := i.chain.All()
	if err != nil {
		return err
	}

	for _, entry := range e {
		handler(entry)
	}

	return nil
}

func (i *Iterator) Dir(handler EachDirHandler) error {
	d, err := i.chain.Dirs()
	if err != nil {
		return err
	}

	for _, dir := range d {
		handler(dir)
	}

	return nil
}

func (i *Iterator) File(handler EachFileHandler) error {
	f, err := i.chain.Files()
	if err != nil {
		return err
	}

	for _, file := range f {
		handler(file)
	}

	return nil
}
