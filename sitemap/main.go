package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"stark/gophercises/linkparser"
)

// Generate generates the sitemap of the given url
func main() {
	url := flag.String("url", "https://www.calhoun.io", "url of website to generate sitemap")
	flag.Parse()

	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}

	// parser accepts type io.Reader and resp.Body is of type io.ReadCloser
	// and ReadCloser is an interface made of two interfaces - io.Reader and io.Closer
	links, err := linkparser.NewParser(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {
		resp, err := http.Get(link.Href)
		if err == nil {
			res, err := linkparser.NewParser(resp.Body)
			if err == nil {
				links = append(links, res...)
			}
		}
	}

	for _, link := range links {
		fmt.Println(link.Href)
	}
}
