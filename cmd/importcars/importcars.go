package main

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	log.Print("importing cars")

	_, err := git.PlainClone("/tmp/cars", false, &git.CloneOptions{
		URL:      "https://github.com/abhionlyone/us-car-models-data",
		Progress: os.Stdout,
	})

	if err != nil {
		log.Fatal("unable to clone cars repo")
	}

}
