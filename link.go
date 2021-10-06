package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var allLinks []Link // make([]Link, 0)

func dfs(n *html.Node, isParentLink bool) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		isParentLink = true
	}
	currentText := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		currentText += dfs(c, isParentLink)
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				link := Link{attr.Val, currentText}
				allLinks = append(allLinks, link)
				break
			}
		}
	} else if n.Type == html.TextNode {
		currentText += n.Data
	}
	return strings.Join(strings.Fields(currentText), " ") // string.Feilds remove all leading and tailling spaces of the string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, false)
	return allLinks, nil
}
