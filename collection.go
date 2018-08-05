package filesystem

type Collection struct {
	source Source
}

func (c *Collection) Each() *Iterator {
	return &Iterator{
		collection: c,
	}
}

func (c *Collection) Exists(path string) bool {
	return c.source.exists(path)
}

func (c *Collection) IsDir(path string) bool {
	return c.source.isDir(path)
}

func (c *Collection) IsFile(path string) bool {
	return c.source.isFile(path)
}

func (c *Collection) Dir(path string) (*Dir, error) {
	if !c.Exists(path) {
		return nil, &ErrNotExists{path: path}
	}

	if !c.IsDir(path) {
		return nil, &ErrNotDir{path: path}
	}

	e, err := c.source.entry(path)
	if err != nil {
		return nil, err
	}

	dir, _ := e.(*Dir)
	return dir, nil
}

func (c *Collection) File(path string) (*File, error) {
	if !c.Exists(path) {
		return nil, &ErrNotExists{path: path}
	}

	if !c.IsFile(path) {
		return nil, &ErrNotFile{path: path}
	}

	e, err := c.source.entry(path)
	if err != nil {
		return nil, err
	}

	file, _ := e.(*File)
	return file, nil
}
