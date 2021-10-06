package link

import (
	"io"
)

type Link struct {
	Href string
	Link string
}

func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
