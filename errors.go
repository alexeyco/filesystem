package filesystem

import (
	"fmt"
	"os"
)

type ErrAlreadyExist struct {
	path string
}

func (e *ErrAlreadyExist) Error() string {
	return fmt.Sprintf("path %s already shouldBeExist", e.path)
}

type ErrNotExist struct {
	path string
}

func (e *ErrNotExist) Error() string {
	return fmt.Sprintf("path %s not shouldBeExist", e.path)
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

type ErrNotInRoot struct {
	path string
}

func (e *ErrNotInRoot) Error() string {
	return fmt.Sprintf("path %s is outside of root directory", e.path)
}

func checkNotExist(path string, err error) error {
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
