package filesystem

var extEditableDefault = []string{
	"txt",
}

type FileEditable struct {
	*File
}

func (f *FileEditable) Type() int {
	return FileTypeEditable
}
