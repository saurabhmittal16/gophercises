package linkparser

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Link contains info about an 'a' tag
type Link struct {
	Href string
	Text string
}

func f(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				foundLink := Link{Href: a.Val, Text: getText(n)}
				fmt.Printf("%+v\n", foundLink)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c)
	}
}

func getText(n *html.Node) string {
	if n != nil {
		if n.Type == html.TextNode {
			return n.Data
		} else if n.Type == html.ElementNode {
			var ans string
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				ans = ans + " " + strings.TrimSpace(getText(c))
			}
			return ans
		}
	}

	return ""
}

// NewParser parses the provided html and finds the anchor tags
func NewParser(r *os.File) {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	f(doc)
}
