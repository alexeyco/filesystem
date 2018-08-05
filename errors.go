package filesystem

import (
	"fmt"
	"os"
)

type ErrNotExist struct {
	path string
}

func (e *ErrNotExist) Error() string {
	return fmt.Sprintf("path %s not exist", e.path)
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

func checkNotNotExist(path string, err error) error {
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		return &ErrNotExist{
			path: path,
		}
	}

	return err
}
