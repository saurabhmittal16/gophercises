package main

import (
	"flag"
	"fmt"
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

	links, err := linkparser.NewParser(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(links)
}
