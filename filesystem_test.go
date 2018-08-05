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

// TestFs_Abs check Fs absolute path
func TestFs_Abs(t *testing.T) {
	tmp, err := mkTmpDir()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(tmp)

	if err := mkDir(tmp, "root"); err != nil {
		t.Error("Can't create temporary file", err)
		return
	}

	root := filepath.Join(tmp, "root")
	fs, err := Root(root)
	if err != nil {
		t.Error(err)
	}

	if root != fs.Abs() {
		t.Errorf("Fs must be %s, not %s", root, fs.Abs())
	}
}

// TestFs_List check Fs list
func TestFs_List(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkDir(fs.Abs(), "bar")
	mkFile(fs.Abs(), "baz.txt")

	dirs, err := fs.List()
	if err != nil {
		t.Error("Can't list root directory", err)
		return
	}

	l := len(dirs)
	if l != 3 {
		t.Errorf("List length must be 3, not %d", l)
	}
}

// TestFs_Dirs check Fs directories list
func TestFs_Dirs(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkDir(fs.Abs(), "bar")
	mkFile(fs.Abs(), "baz.txt")

	dirs, err := fs.Dirs()
	if err != nil {
		t.Error("Can't list root directory", err)
		return
	}

	l := len(dirs)
	if l != 2 {
		t.Errorf("Dirs count must be 2, not %d", l)
	}
}

// TestFs_DirNotFound check Fs.Dir() error
func TestFs_DirNotFound(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkDir(fs.Abs(), "bar")
	mkFile(fs.Abs(), "baz.txt")

	dir, err := fs.Dir("fiz")
	if err == nil {
		t.Error("Must be error")
	} else {
		if _, ok := err.(*ErrDirNotFound); !ok {
			t.Error("Error must have *ErrDirNotFound type", err)
		}
	}

	if dir != nil {
		t.Error("Directory must be null, not", dir)
	}

	fs, err = Root(fs.Abs())
	if err != nil {
		t.Error(err)
		return
	}

	os.RemoveAll(filepath.Join(fs.Abs()))
	_, err = fs.Dir("foo")
	if err == nil {
		t.Error("Must be error")
	} else {
		if _, ok := err.(*ErrDirNotFound); !ok {
			t.Error("Error must have *ErrDirNotFound type", err)
		}
	}
}

// TestFs_DirOk check Fs.Dir()
func TestFs_DirOk(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkDir(fs.Abs(), "bar")
	mkFile(fs.Abs(), "baz.txt")

	dir, err := fs.Dir("foo")
	if err != nil {
		t.Error("There must be a \"foo\" directory", err)
		return
	}

	if dir.Name() != "foo" {
		t.Errorf("Directory name must be \"foo\", not %s", dir.Name())
	}

	if dir.Path() != "foo" {
		t.Errorf("Directory path must be \"foo\", not %s", dir.Name())
	}
}

// TestFs_Files check Fs files list
func TestFs_Files(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkFile(fs.Abs(), "bar.txt")
	mkFile(fs.Abs(), "baz.txt")

	files, err := fs.Files()
	if err != nil {
		t.Error("Can't list root files", err)
		return
	}

	l := len(files)
	if l != 2 {
		t.Errorf("Files count must be 2, not %d", l)
	}
}

// TestFs_FileNotFound check Fs.File() error
func TestFs_FileNotFound(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkDir(fs.Abs(), "bar")
	mkFile(fs.Abs(), "baz.txt")

	file, err := fs.File("fiz.txt")
	if err == nil {
		t.Error("Must be errors")
	} else {
		if _, ok := err.(*ErrFileNotFound); !ok {
			t.Error("Error must have *ErrFileNotFound type", err)
		}
	}

	if file != nil {
		t.Error("File must be null, not", file)
	}

	fs, err = Root(fs.Abs())
	if err != nil {
		t.Error(err)
		return
	}

	os.RemoveAll(fs.Abs())

	_, err = fs.File("bar.txt")
	if err == nil {
		t.Error("Must be error")
	} else {
		if _, ok := err.(*ErrDirNotFound); !ok {
			t.Error("Error must have *ErrFileNotFound type", err)
		}
	}
}

// TestFs_FileOk check Fs.File()
func TestFs_FileOk(t *testing.T) {
	fs, err := getFsTmp()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(fs.Abs())

	mkDir(fs.Abs(), "foo")
	mkDir(fs.Abs(), "bar")
	mkFile(fs.Abs(), "baz.txt")

	file, err := fs.File("baz.txt")
	if err != nil {
		t.Error("There must be a \"baz.txt\" file", err)
		return
	}

	if file.Name() != "baz.txt" {
		t.Errorf("Directory name must be \"baz.txt\", not %s", file.Name())
	}

	if file.Path() != "baz.txt" {
		t.Errorf("Directory path must be \"baz.txt\", not %s", file.Name())
	}
}

// TestRoot_NotFound check wrong root error
func TestRoot_NotFound(t *testing.T) {
	_, err := Root("./wrong/directory")
	if err == nil {
		t.Error("Must be error")
		return
	}

	if _, ok := err.(*ErrDirNotFound); !ok {
		t.Error("Must be ErrNotFound error", err)
	}
}

// TestRoot_MustBeDirectory check wrong root error
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
	tmp, err := mkTmpDir()
	if err != nil {
		t.Error("Can't create temporary Fs root", err)
		return
	}
	defer os.RemoveAll(tmp)

	if err := mkDir(tmp, "root"); err != nil {
		t.Error("Can't create temporary file", err)
		return
	}

	_, err = Root(filepath.Join(tmp, "root"))
	if err != nil {
		t.Error(err)
	}
}
