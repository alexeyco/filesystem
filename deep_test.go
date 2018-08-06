package filesystem

import (
	"testing"
	"os"
)

func TestSeekerDeep_Root(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	mkDir(root.Root(), "foo/bar/baz")

	dirs := 0
	err = root.Read(Deep()).Each().Dir(func(dir *Dir) {
		dirs++
	})

	if dirs != 3 {
		t.Errorf("Should be %d directories, not %d", 3, dirs)
	}
}
