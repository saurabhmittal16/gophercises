package main

import (
	"flag"
	"log"
	"os"
	"stark/gophercises/linkparser"
)

func main() {
	filename := flag.String("file", "ex1.html", "select file to run parser")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	linkparser.NewParser(file)
}
