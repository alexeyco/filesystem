package main

import (
	"fmt"
	"log"

	"github.com/alexeyco/filesystem"
)

func main() {
	fs, err := filesystem.Root("../data")
	if err != nil {
		log.Fatalln(err)
	}

	// Source of ../data/foo
	collection := fs.Read(filesystem.In("foo"))

	err = collection.Each().Dir(func(dir *filesystem.Dir) {
		fmt.Println("Dir:  ", dir.Name())
	})

	if err != nil {
		log.Fatalln(err)
	}

	err = collection.Each().File(func(file *filesystem.File) {
		fmt.Println("File: ", file.Name())
	})

	if err != nil {
		log.Fatalln(err)
	}

	// Get file ../data/foo/fizz/bar.txt info
	file, err := collection.File("fizz/buzz.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()
	fmt.Println("File: ", file.Name())
}
