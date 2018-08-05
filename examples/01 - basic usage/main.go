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

	err = fs.Each().Entry(func(entry filesystem.Entry) {
		fmt.Println(entry.Name())
	})

	if err != nil {
		log.Fatalln(err)
	}
}
