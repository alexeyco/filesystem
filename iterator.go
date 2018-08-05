package filesystem

type Iterator struct {
	seeker Seeker
}

func (i *Iterator) Entry(handler EachEntryHandler) error {
	return i.seeker.each(handler)
}

func (i *Iterator) Dir(handler EachDirHandler) error {
	return i.seeker.each(func(entry Entry) {
		if entry.IsDir() {
			e, _ := entry.(*Dir)
			handler(e)
		}
	})
}

func (i *Iterator) File(handler EachFileHandler) error {
	return i.seeker.each(func(entry Entry) {
		if entry.IsFile() {
			f, _ := entry.(*File)
			handler(f)
		}
	})
}
