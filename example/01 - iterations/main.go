package main

import (
	"fmt"
	"log"

	"github.com/alexeyco/filesystem"
)

func main() {
	fs, err := filesystem.Root("../testdata")
	if err != nil {
		log.Fatalln(err)
	}

	dirs, err := fs.Dirs()
	if err != nil {
		log.Fatalln(err)
	}

	files, err := fs.Files()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Contents of ", fs.Root())
	dirs.Each(func(dir *filesystem.Dir) {
		fmt.Println("Dir:  ", dir.Name())
	})

	files.Each(func(file *filesystem.File) {
		fmt.Println("File: ", file.Name())
	})
}
