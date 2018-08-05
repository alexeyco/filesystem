package filesystem

import (
	"os"
	"testing"
)

// TestFiles_Each check files iteration
func TestFiles_Each(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkFile(fs.Abs(), "bar.txt")

	files, err := fs.Files()
	if err != nil {
		t.Error("Must be no errors", err)
		return
	}

	parent := fs.root
	files.Each(func(file *File) {
		if file.IsDir() {
			t.Error("File can't be dir")
		}

		if file.Name() != "bar.txt" {
			t.Errorf("File must have name \"bar.txt\", not %s", file.Name())
		}

		if file.Parent() != parent {
			t.Errorf("Parent must be \"%s\", not \"%s\"", parent.Path(), file.Path())
		}
	})
}
