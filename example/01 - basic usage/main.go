package main

import (
	"fmt"
	"log"

	"github.com/alexeyco/filesystem"
)

func main() {
	fs, err := filesystem.Root("./testdata")
	if err != nil {
		log.Fatalln(err)
	}

	paths, err := fs.List()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Contents of ", fs.)
	for _, p := range paths {

	}
}
