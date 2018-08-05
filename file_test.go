package filesystem

import (
	"os"
	"path/filepath"
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

// TestFile_Rename checks file rename
func TestFile_Rename(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkFile(fs.Abs(), "foo.txt")

	file, err := fs.File("foo.txt")
	if err != nil {
		t.Error(err)
		return
	}

	if err := file.Rename("bar.txt"); err != nil {
		t.Error(err)
		return
	}

	files, err := fs.Files()
	if err != nil {
		t.Error(err)
		return
	}

	l := len(files)
	if l != 1 {
		t.Errorf("Must be 1 file, not %d", l)
		return
	}

	_, err = fs.File("bar.txt")
	if err != nil {
		t.Error(err)
	}
}

// TestFile_Move checks file move
func TestFile_Move(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkDir(fs.Abs(), "bar")
	mkFile(filepath.Join(fs.Abs(), "foo"), "baz.txt")

	foo, err := fs.Dir("foo")
	if err != nil {
		t.Error(err)
		return
	}

	bar, err := fs.Dir("foo")
	if err != nil {
		t.Error(err)
		return
	}

	baz, err := foo.File("baz.txt")
	if err != nil {
		t.Error(err)
		return
	}

	if err := baz.Move(bar); err != nil {
		t.Error(err)
		return
	}

	if baz.Parent() != bar {
		t.Errorf("Parent must be \"%s\", not \"%s\"", bar.Name(), baz.Parent().Name())
		return
	}

	p := filepath.Join(bar.Path(), baz.Name())
	if baz.Path() != p {
		t.Errorf("Path must be \"%s\", not \"%s\"", p, baz.Path())
		return
	}

	_, err = bar.File("baz.txt")
	if err != nil {
		t.Error(err)
		return
	}
}

// TestFile_Remove checks file remove
func TestFile_Remove(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkFile(fs.Abs(), "foo.txt")
	mkFile(fs.Abs(), "bar.txt")

	file, err := fs.File("bar.txt")
	if err != nil {
		t.Error(err)
		return
	}

	if err := file.Remove(); err != nil {
		t.Error(err)
		return
	}

	files, err := fs.Files()
	if err != nil {
		t.Error(err)
		return
	}

	l := len(files)
	if l != 1 {
		t.Errorf("Must be 1 file, not %d", l)
	}
}