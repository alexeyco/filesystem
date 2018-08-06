package filesystem

// Iterator files iterator
type Iterator struct {
	seeker Seeker
}

// Entry iterates through files and directories
func (i *Iterator) Entry(handler EachEntryHandler) error {
	return i.seeker.each(handler)
}

// Dir iterates through directories
func (i *Iterator) Dir(handler EachDirHandler) error {
	return i.Entry(func(entry Entry) {
		if entry.IsDir() {
			e, _ := entry.(*Dir)
			handler(e)
		}
	})
}

// File iterates through files
func (i *Iterator) File(handler EachFileHandler) error {
	return i.Entry(func(entry Entry) {
		if entry.IsFile() {
			f, _ := entry.(*File)
			handler(f)
		}
	})
}
