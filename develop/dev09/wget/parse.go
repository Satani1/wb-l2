package wget

import (
	"golang.org/x/net/html"
	"io"
	"net/url"
)

// Link represents html link tag
type Link struct {
	href *url.URL
}

// ParseHTML parse given html file and returns slice of links
func ParseHTML(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	links := make([]Link, 0)

	var parseNode func(node *html.Node)
	parseNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			u, err := ParseHref(n.Attr)
			if err != nil {
				return
			}
			links = append(links, Link{href: u})
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseNode(c)
		}
	}
	parseNode(doc)

	return links, nil
}

// ParseHref extract href attribute from link tag
func ParseHref(attrs []html.Attribute) (*url.URL, error) {
	var href string

	for _, a := range attrs {
		if a.Key == "href" {
			href = a.Val
			break
		}
	}
	u, err := url.Parse(href)
	if err != nil {
		return nil, err
	}
	return u, nil
}
