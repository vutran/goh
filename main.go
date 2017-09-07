package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

func MakeDirectory(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		fi, _ := os.Stat(path.Dir(dir))
		mkErr := os.Mkdir(dir, fi.Mode())
		if mkErr != nil {
			log.Fatal("Error trying to create directory")
		}
	}
}

func main() {
	mkdir := flag.Bool("m", false, "Create the directory if necessary")
	flag.Parse()
	name := flag.Arg(0)

	fullPath := path.Join(os.Getenv("GOPATH"), "src", "github.com", name)

	if *mkdir {
		MakeDirectory(fullPath)
	}

	fmt.Print(fullPath)
}
