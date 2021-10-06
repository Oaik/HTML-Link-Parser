package main

import (
	"fmt"
	"link-parser"
	"strings"
)

var html = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(html)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	len := len(links)
	for i := 0; i < len; i++ {
		fmt.Printf("Href: %v\nText: %v\n", links[i].Href, links[i].Text)
	}
}
