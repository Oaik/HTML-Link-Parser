package main

import (
	"flag"
	"fmt"
	"link-parser"
	"log"
	"os"
)

func main() {
	fptr := flag.String("htmltag", "index.html", "file path to read from")
	flag.Parse()
	r, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	for _, link := range links {
		fmt.Printf("Href: %v\nText: %v\n", link.Href, link.Text)
	}
}
