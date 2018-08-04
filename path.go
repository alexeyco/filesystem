package filesystem

type Path interface {
	IsFile() bool
	IsDir() bool
	Parent() *Dir
	Name() string
	Path() string
	Rename() error
	Move(dir *Dir) error
	Remove() error
}
