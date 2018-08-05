package filesystem

import (
	"os"
	"testing"
)

// TestDirs_Each check directories iteration
func TestDirs_Each(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkFile(fs.Abs(), "bar.txt")

	dirs, err := fs.Dirs()
	if err != nil {
		t.Error("Must be no errors", err)
		return
	}

	parent := fs.root
	dirs.Each(func(dir *Dir) {
		if dir.IsFile() {
			t.Error("Dir can't be file")
		}

		if dir.Name() != "foo" {
			t.Errorf("Directory must have name \"foo\", not %s", dir.Name())
		}

		if dir.Parent() != parent {
			t.Errorf("Parent must be \"%s\", not \"%s\"", parent.Path(), dir.Path())
		}
	})
}
