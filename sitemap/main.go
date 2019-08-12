package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"stark/gophercises/linkparser"
)

const xmlnamespace = "http://www.sitemaps.org/schemas/sitemap/0.9"

type empty struct{}

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

// Generate generates the sitemap of the given url
func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "url of website to generate sitemap")
	maxDepth := flag.Int("depth", 3, "maximum depth for link checking")
	flag.Parse()

	// get pages
	pages := bfs(*urlFlag, *maxDepth)

	toXML := urlset{
		Xmlns: xmlnamespace,
	}

	for _, page := range pages {
		toXML.Urls = append(toXML.Urls, loc{page})
	}

	// print generated sitemap XML to stdout
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", " ")
	if err := enc.Encode(toXML); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func bfs(url string, depth int) []string {
	// visited is a set of visited urls
	visited := make(map[string]empty)

	// q is a set that will contain the current level of tree
	q := make(map[string]empty)

	// nq contains the next level of tree
	nq := map[string]empty{
		url: empty{},
	}

	// iterate for depth times
	for i := 0; i <= depth; i++ {
		// set q as nq and initialise nq
		q, nq = nq, make(map[string]empty)

		if len(q) == 0 {
			break
		}

		// for every url in current level
		// check if it has been visited, if not then get it's urls
		for url := range q {
			if _, ok := visited[url]; ok {
				continue
			}
			visited[url] = empty{}

			for _, link := range get(url) {
				if _, ok := visited[link]; !ok {
					nq[link] = empty{}
				}
			}
		}
	}

	ret := make([]string, 0, len(visited))
	for key := range visited {
		ret = append(ret, key)
	}

	return ret
}

// get sends a request to urlString and generates the base URL and the parsed links
func get(urlString string) []string {
	resp, err := http.Get(urlString)
	if err != nil {
		log.Fatal(err)
	}

	// get request URL and base URL from the response body
	rURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: rURL.Scheme,
		Host:   rURL.Host,
	}
	base := baseURL.String()

	// get hrefs from the response body and filter them
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

// hrefs parses the reader to find hrefs and returns clean links
func hrefs(r io.Reader, base string) []string {
	// slice to store final urls
	var ret []string

	links, err := linkparser.NewParser(r)
	if err != nil {
		log.Fatal(err)
	}

	// clean links to contain either /some-url or complete urls like https://gophercises.com
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}

	return ret
}

// filter filters the slice of strings by invoking the keepFn method
func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}

	return ret
}

// closure that returns a function which uses the passed value
func withPrefix(base string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, base)
	}
}
