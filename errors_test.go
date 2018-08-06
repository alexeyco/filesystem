package filesystem

import (
	"strings"
	"testing"
)

func TestErrors(t *testing.T) {
	path := "foobar"

	errors := []error{
		&ErrAlreadyExist{path: path},
		&ErrNotExist{path: path},
		&ErrNotDir{path: path},
		&ErrNotFile{path: path},
		&ErrNotInRoot{path: path},
	}

	for _, err := range errors {
		if !strings.Contains(err.Error(), path) {
			t.Error("Error text should contain \"foobar\"", err)
		}
	}
}
