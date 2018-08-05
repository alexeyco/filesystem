package main

import (
	"fmt"
	"log"

	"github.com/alexeyco/filesystem"
)

func main() {
	root, err := filesystem.Root("../data")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("List root directory", root.Root(), "contents")
	err = root.Each().Dir(func(dir *filesystem.Dir) {
		fmt.Println("Dir:  ", dir.Name())
	})

	err = root.Each().File(func(file *filesystem.File) {
		fmt.Println("File: ", file.Name())
	})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()

	foo := root.Read(filesystem.In("foo"))

	fmt.Println("List", foo.Root(), "contents")
	err = foo.Each().Dir(func(dir *filesystem.Dir) {
		fmt.Println("Dir:  ", dir.Name())
	})

	if err != nil {
		log.Fatalln(err)
	}

	err = foo.Each().File(func(file *filesystem.File) {
		fmt.Println("File: ", file.Name())
	})

	if err != nil {
		log.Fatalln(err)
	}
}
