package filesystem

import (
	"os"
	"testing"
)

func TestIterator(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	mkDir(root.Root(), "foo")
	mkDir(root.Root(), "bar")
	mkFile(root.Root(), "foo.txt")
	mkFile(root.Root(), "bar.txt")
	mkFile(root.Root(), "fiz.txt")
	mkFile(root.Root(), "buz.txt")

	dirs := 0
	files := 0

	err = root.Each().Dir(func(dir *Dir) {
		dirs++
	})

	if err != nil {
		t.Error(err)
	}

	err = root.Each().File(func(file *File) {
		files++
	})

	if err != nil {
		t.Error(err)
	}

	if dirs != 2 {
		t.Errorf("Must be %d directories, not %d", 2, dirs)
	}

	if files != 4 {
		t.Errorf("Must be %d files, not %d", 4, files)
	}
}
