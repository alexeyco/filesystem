package filesystem

// Seeker file seeker interface
type Seeker interface {
	SetRoot(string)
	Root() string
	Each() *Iterator
	Exist(path string) bool
	IsDir(path string) bool
	IsFile(path string) bool
	Dir(path string) (*Dir, error)
	File(path string) (*File, error)
	Mkdir(path string) error
	Move(source, dest string) error
	Remove(path string) error
	each(handler EachEntryHandler) error
}
