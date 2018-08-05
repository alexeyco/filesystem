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
	_, err = Root(filepath.Join(dir, "wrong-directory"))
	if err == nil {
		t.Error("Should be error")
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

	if root.Exist("../..") {
		t.Error("there should be nothing outside the root")
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

	_, err = root.Dir("bar.txt")
	if err == nil {
		t.Error("Should be errors")
	}

	if _, ok := err.(*ErrNotDir); !ok {
		t.Error("Should be *ErrNotDir type", err)
	}

	_, err = root.Dir("fizz")
	if err == nil {
		t.Error("Should be errors")
	}

	if _, ok := err.(*ErrNotExist); !ok {
		t.Error("Should be *ErrNotExist type", err)
	}

	_, err = root.Dir("../..")
	if err == nil {
		t.Error("Should be error")
		return
	}

	if _, ok := err.(*ErrNotInRoot); !ok {
		t.Error("Should be *ErrNotInRoot type", err)
		return
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

	_, err = root.File("foo")
	if err == nil {
		t.Error("Should be errors")
	}

	if _, ok := err.(*ErrNotFile); !ok {
		t.Error("Should be *ErrNotFile type", err)
	}

	_, err = root.File("fizz.txt")
	if err == nil {
		t.Error("Should be errors")
	}

	if _, ok := err.(*ErrNotExist); !ok {
		t.Error("Should be *ErrNotExist type", err)
	}

	_, err = root.File("../..")
	if err == nil {
		t.Error("Should be error")
		return
	}

	if _, ok := err.(*ErrNotInRoot); !ok {
		t.Error("Should be *ErrNotInRoot type", err)
		return
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

	err = root.Mkdir("../../foo")
	if err == nil {
		t.Error("Should be error")
		return
	}

	if _, ok := err.(*ErrNotInRoot); !ok {
		t.Error("Should be *ErrNotInRoot type", err)
		return
	}
}

func TestFs_Move(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	mkDir(root.Root(), "foo/bar/baz")
	mkFile(root.Root(), "bar")
	mkFile(root.Root(), "bar.txt")

	if err := root.Move("bar", "foo/bar/baz/bar"); err != nil {
		t.Error(err)
	}

	if err := root.Move("bar.txt", "foo/bar/baz/bar.txt"); err != nil {
		t.Error(err)
	}

	if root.Exist("bar") {
		t.Errorf("\"%s\" should not exist", "bar")
	}

	if root.Exist("bar.txt") {
		t.Errorf("\"%s\" should not exist", "bar.txt")
	}

	if !root.Exist("foo/bar/baz/bar") {
		t.Errorf("\"%s\" should not exist", "foo/bar/baz/bar")
	}

	if !root.Exist("foo/bar/baz/bar.txt") {
		t.Errorf("\"%s\" should not exist", "foo/bar/baz/bar.txt")
	}

	err = root.Move("foo/bar/baz/bar.txt", "foo/bar/baz/bar")
	if err == nil {
		t.Error("Should be error")
		return
	}

	if _, ok := err.(*ErrAlreadyExist); !ok {
		t.Error("Should be *ErrAlreadyExist type", err)
		return
	}

	err = root.Move("bar", "foo/fizz")
	if err == nil {
		t.Error("Should be error")
		return
	}

	if _, ok := err.(*ErrNotExist); !ok {
		t.Error("Should be *ErrNotExist type", err)
		return
	}

	err = root.Move("../..", "foo/fizz")
	if err == nil {
		t.Error("Should be error")
		return
	}

	if _, ok := err.(*ErrNotInRoot); !ok {
		t.Error("Should be *ErrNotInRoot type", err)
		return
	}

	err = root.Move("foo", "../..")
	if err == nil {
		t.Error("Should be error")
		return
	}

	if _, ok := err.(*ErrNotInRoot); !ok {
		t.Error("Should be *ErrNotInRoot type", err)
		return
	}
}

func TestFs_Remove(t *testing.T) {
	root, err := getTmpRoot()
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(root.Root())

	mkDir(root.Root(), "foo")
	mkFile(root.Root(), "bar.txt")

	if err := root.Remove("foo"); err != nil {
		t.Error(err)
	}

	if root.Exist("foo") {
		t.Errorf("\"%s\" should not exist", "foo")
	}

	if err := root.Remove("bar.txt"); err != nil {
		t.Error(err)
	}

	if root.Exist("bar.txt") {
		t.Errorf("\"%s\" should not exist", "bar.txt")
	}

	err = root.Remove("bar.txt")
	if err == nil {
		t.Error("Should be error")
		return
	}

	if _, ok := err.(*ErrNotExist); !ok {
		t.Error("Should be *ErrNotExist type", err)
		return
	}

	err = root.Remove("../..")
	if err == nil {
		t.Error("Should be error")
		return
	}

	if _, ok := err.(*ErrNotInRoot); !ok {
		t.Error("Should be *ErrNotInRoot type", err)
		return
	}
}
