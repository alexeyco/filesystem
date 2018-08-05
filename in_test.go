package filesystem

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSeekerIn_Root(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	foo := root.Read(In("foo"))

	r := filepath.Join(root.Root(), "foo")
	if foo.Root() != r {
		t.Errorf("Root should be \"%s\", not \"%s\"", r, foo.Root())
	}
}
