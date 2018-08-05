package filesystem

type Collection struct {
	source Source
}

func (c *Collection) Each() *Iterator {
	return &Iterator{
		collection: c,
	}
}
