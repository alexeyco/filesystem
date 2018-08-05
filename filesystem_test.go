package filesystem

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func getFsTmp() (*Fs, error) {
	dir, err := mkTmpDir()
	if err != nil {
		return nil, err
	}

	return Root(dir)
}

func mkTmpDir() (string, error) {
	return ioutil.TempDir("", "fs-test")
}

func mkDir(root, name string) error {
	return os.MkdirAll(filepath.Join(root, name), os.ModePerm)
}

func mkFile(root, name string) error {
	name = filepath.Join(root, name)
	return ioutil.WriteFile(name, []byte{}, os.ModePerm)
}

// TestRoot_NotExist when trying to open a nonexistent directory, there must be an appropriate error
func TestRoot_NotExist(t *testing.T) {
	_, err := Root("./wrong/directory")
	if err == nil {
		t.Error("Must be error")
		return
	}

	if _, ok := err.(*ErrDirNotFound); !ok {
		t.Error("Must be ErrNotFound error", err)
	}
}

// TestRoot_Ok only a directory can be a root
func TestRoot_MustBeDirectory(t *testing.T) {
	tmp, err := mkTmpDir()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(tmp)

	if err := mkFile(tmp, "root"); err != nil {
		t.Error("Can't create temporary file", err)
		return
	}

	_, err = Root(filepath.Join(tmp, "root"))
	if err == nil {
		t.Error("Root should not be a file")
	}
}

// TestRoot_Ok all requirements are met
func TestRoot_Ok(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())
}
