package link

import (
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	allLinks := []Link{} // make([]Link, 0)
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var dfs func(*html.Node, bool) string
	dfs = func(n *html.Node, isParentLink bool) string {
		if n.Type == html.ElementNode && n.Data == "a" {
			isParentLink = true
		}
		currentText := ""
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			currentText += dfs(c, isParentLink)
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			len := len(n.Attr)
			for i := 0; i < len; i++ {
				if n.Attr[i].Key == "href" {
					var link Link
					link.Href = n.Attr[i].Val
					link.Text = currentText
					allLinks = append(allLinks, link)
				}
			}
		}
		if n.Type == html.TextNode {
			currentText += n.Data
		}
		return currentText
	}
	dfs(doc, false)
	return allLinks, nil
}
