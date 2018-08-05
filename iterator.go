package filesystem

type Iterator struct {
	collection *Collection
}

func (i *Iterator) Entry(handler EachEntryHandler) error {
	e, err := i.collection.source.entries()
	if err != nil {
		return err
	}

	for _, entry := range e {
		handler(entry)
	}

	return nil
}

func (i *Iterator) Dir(handler EachDirHandler) error {
	return i.Entry(func(entry Entry) {
		if entry.IsDir() {
			dir, _ := entry.(*Dir)
			handler(dir)
		}
	})
}

func (i *Iterator) File(handler EachFileHandler) error {
	return i.Entry(func(entry Entry) {
		if !entry.IsDir() {
			file, _ := entry.(*File)
			handler(file)
		}
	})
}
