package filesystem

import (
	"fmt"
	"os"
)

// ErrAlreadyExist file already exist
type ErrAlreadyExist struct {
	path string
}

// Error returns error string
func (e *ErrAlreadyExist) Error() string {
	return fmt.Sprintf("path %s already shouldBeExist", e.path)
}

// ErrNotExist path is not exist
type ErrNotExist struct {
	path string
}

// Error returns error string
func (e *ErrNotExist) Error() string {
	return fmt.Sprintf("path %s not shouldBeExist", e.path)
}

// ErrNotDir path is not directory
type ErrNotDir struct {
	path string
}

// Error returns error string
func (e *ErrNotDir) Error() string {
	return fmt.Sprintf("path %s is not a directory", e.path)
}

// ErrNotFile path is not file
type ErrNotFile struct {
	path string
}

// Error returns error string
func (e *ErrNotFile) Error() string {
	return fmt.Sprintf("path %s is not a file", e.path)
}

// ErrNotInRoot path is outside root directory
type ErrNotInRoot struct {
	path string
}

// Error returns error string
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
