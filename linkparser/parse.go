package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link contains info about an 'a' tag
type Link struct {
	Href string
	Text string
}

func f(n *html.Node) []Link {
	res := []Link{}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				foundLink := Link{Href: a.Val, Text: getText(n)}
				res = append(res, foundLink)
				// fmt.Printf("%+v\n", foundLink)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, f(c)...)
	}

	return res
}

func getText(n *html.Node) string {
	if n != nil {
		if n.Type == html.TextNode {
			return n.Data
		} else if n.Type == html.ElementNode {
			var ans string
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				ans += getText(c)
			}

			return strings.Join(strings.Fields(ans), " ")
		}
	}

	return ""
}

// NewParser parses the provided html and finds the anchor tags
func NewParser(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	res := f(doc)
	return res, nil
}
