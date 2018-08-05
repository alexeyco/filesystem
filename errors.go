package filesystem

import "fmt"

type ErrNotExists struct {
	path string
}

func (e *ErrNotExists) Error() string {
	return fmt.Sprintf("in %s not exists", e.path)
}

type ErrNotDir struct {
	path string
}

func (e *ErrNotDir) Error() string {
	return fmt.Sprintf("path %s is not a directory", e.path)
}

type ErrNotFile struct {
	path string
}

func (e *ErrNotFile) Error() string {
	return fmt.Sprintf("path %s is not a file", e.path)
}
