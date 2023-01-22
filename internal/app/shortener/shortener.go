package shortener

import (
	"strconv"
)

type Shortener struct {
}

func (s *Shortener) Short(link string) string {
	linkLength := len(link)
	return strconv.Itoa(linkLength)
}

func New() *Shortener {
	return &Shortener{}
}
