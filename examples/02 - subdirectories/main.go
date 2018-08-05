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

	collection := fs.Read(filesystem.In("foo"))

	err = collection.Each().Dir(func(dir *filesystem.Dir) {
		fmt.Println("Dir:  ", dir.Name())
	})

	err = collection.Each().File(func(file *filesystem.File) {
		fmt.Println("File: ", file.Name())
	})

	if err != nil {
		log.Fatalln(err)
	}
}
