package filesystem

type FileImage struct {
	*File
}

func (f *FileImage) Type() int {
	return FileTypeImage
}
