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

	err = root.Read(filesystem.Deep()).Each().Entry(func(entry filesystem.Entry) {
		if entry.IsDir() {
			fmt.Println("Dir:  ", entry.Name())
		} else {
			fmt.Println("File: ", entry.Name())
		}
	})

	if err != nil {
		log.Fatalln(err)
	}
}
