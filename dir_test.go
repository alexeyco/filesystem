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

// TestDir_Rename checks directory rename
func TestDir_Rename(t *testing.T) {

}

// TestDir_Move checks directory move
func TestDir_Move(t *testing.T) {

}

// TestDir_Remove checks directory remove
func TestDir_Remove(t *testing.T) {

}

// TestDir_Mkdir checks directory create
func TestDir_Mkdir(t *testing.T) {

}
