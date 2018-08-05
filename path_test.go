package filesystem

import (
	"os"
	"testing"
)

// TestPaths_Each check files iteration
func TestPaths_Each(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkFile(fs.Abs(), "bar.txt")

	list, err := fs.List()
	if err != nil {
		t.Error("Must be no errors", err)
		return
	}

	parent := fs.root
	list.Each(func(path Path) {
		if path.IsDir() {
			if path.Name() != "foo" {
				t.Errorf("Directory must have name \"foo\", not %s", path.Name())
			}
		}

		if path.IsFile() {
			if path.Name() != "bar.txt" {
				t.Errorf("File must have name \"bar.txt\", not %s", path.Name())
			}
		}

		if path.Parent() != parent {
			t.Errorf("Parent must be \"%s\", not \"%s\"", parent.Path(), path.Path())
		}
	})
}
