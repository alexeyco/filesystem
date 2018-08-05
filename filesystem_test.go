package filesystem

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func getTmpRoot() (*Fs, error) {
	dir, err := mkTmpDir()
	if err != nil {
		return nil, err
	}

	return Root(dir)
}

func mkTmpDir() (string, error) {
	return ioutil.TempDir("", "filepath-test")
}

func mkDir(root, name string) error {
	return os.MkdirAll(filepath.Join(root, name), os.ModePerm)
}

func mkFile(root, name string) error {
	name = filepath.Join(root, name)
	return ioutil.WriteFile(name, []byte{}, os.ModePerm)
}

func TestRoot(t *testing.T) {
	dir, err := mkTmpDir()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(dir)

	// An attempt to specify a non-existent directory should cause an error
	if _, err = Root(filepath.Join(dir, "wrong-directory")); err == nil {
		t.Error("Must be error")
		return
	}

	if _, ok := err.(*ErrNotExist); !ok {
		t.Error("Should be *ErrNotExist type", err)
		return
	}

	// If directory correct, should be ok
	if _, err = Root(dir); err != nil {
		t.Error(err)
	}
}

func TestFs_Root(t *testing.T) {
	dir, err := mkTmpDir()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(dir)

	root, err := Root(dir)
	if err != nil {
		t.Error(err)
		return
	}

	if root.Root() != dir {
		t.Errorf("Incorrect Fs root: should be \"%s\", not \"%s\"", dir, root.Root())
	}
}

func TestFs_Exist(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	if root.Exist("foo") {
		t.Errorf("\"%s\" should not exist", "foo")
	}

	if root.Exist("foo.txt") {
		t.Errorf("\"%s\" should not exist", "foo.txt")
	}

	mkDir(root.Root(), "foo")
	mkFile(root.Root(), "foo.txt")

	if !root.Exist("foo") {
		t.Errorf("\"%s\" should exist", "foo")
	}

	if !root.Exist("foo.txt") {
		t.Errorf("\"%s\" should exist", "foo.txt")
	}
}

func TestFs_IsDir(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	if root.IsDir("foo") {
		t.Errorf("\"%s\" should not exist", "foo")
	}

	mkDir(root.Root(), "foo")
	mkFile(root.Root(), "foo.txt")

	if !root.IsDir("foo") {
		t.Errorf("\"%s\" should be directory", "foo")
	}

	if root.IsDir("foo.txt") {
		t.Errorf("\"%s\" should not be directory", "foo.txt")
	}
}

func TestFs_IsFile(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	if root.IsDir("foo.txt") {
		t.Errorf("\"%s\" should not exist", "foo")
	}

	mkDir(root.Root(), "foo")
	mkFile(root.Root(), "foo.txt")

	if !root.IsFile("foo.txt") {
		t.Errorf("\"%s\" should be file", "foo.txt")
	}

	if root.IsFile("foo") {
		t.Errorf("\"%s\" should not be directory", "foo")
	}
}

func TestFs_Dir(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	mkDir(root.Root(), "foo")
	mkFile(root.Root(), "bar.txt")

	if _, err = root.Dir("foo"); err != nil {
		t.Error(err)
	}

	if _, err = root.Dir("bar.txt"); err == nil {
		t.Error("Should be errors")
	}

	if _, ok := err.(*ErrNotDir); !ok {
		t.Error("Should be *ErrNotDir type", err)
	}

	if _, err = root.Dir("fizz"); err == nil {
		t.Error("Should be errors")
	}

	if _, ok := err.(*ErrNotExist); !ok {
		t.Error("Should be *ErrNotExist type", err)
	}
}

func TestFs_File(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	mkDir(root.Root(), "foo")
	mkFile(root.Root(), "bar.txt")

	if _, err = root.File("bar.txt"); err != nil {
		t.Error(err)
	}

	if _, err = root.File("foo"); err == nil {
		t.Error("Should be errors")
	}

	if _, ok := err.(*ErrNotFile); !ok {
		t.Error("Should be *ErrNotFile type", err)
	}

	if _, err = root.Dir("fizz.txt"); err == nil {
		t.Error("Should be errors")
	}

	if _, ok := err.(*ErrNotExist); !ok {
		t.Error("Should be *ErrNotExist type", err)
	}
}

func TestFs_Mkdir(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	if err := root.Mkdir("foo"); err != nil {
		t.Error(err)
	}

	if !root.IsDir("foo") {
		t.Errorf("\"%s\" should not be directory", "foo")
	}
}

func TestFs_Move(t *testing.T) {

}

func TestFs_Remove(t *testing.T) {

}
