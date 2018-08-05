package filesystem

type Iterator struct {
	seeker Seeker
}

func (i *Iterator) Entry(handler EachEntryHandler) error {
	return nil
}

func (i *Iterator) Dir(handler EachDirHandler) error {
	return nil
}

func (i *Iterator) File(handler EachFileHandler) error {
	return nil
}
