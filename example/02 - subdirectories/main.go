package main

import (
	"log"

	"fmt"
	"github.com/alexeyco/filesystem"
)

func main() {
	fs, err := filesystem.Root("../testdata")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = fs.Dir("nonexistent-directory")
	if err != nil {
		if _, ok := err.(*filesystem.ErrDirNotFound); !ok {
			log.Fatalln("Something went wrong", err)
		}
	}

	d, err := fs.Dir("foo")
	if err != nil {
		log.Fatalln(err)
	}

	dirs, err := d.Dirs()
	if err != nil {
		log.Fatalln(err)
	}

	files, err := d.Files()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Contents of ", d.Path())
	dirs.Each(func(dir *filesystem.Dir) {
		fmt.Println("Dir:  ", dir.Name())
	})

	files.Each(func(file *filesystem.File) {
		fmt.Println("File: ", file.Name())
	})
}
