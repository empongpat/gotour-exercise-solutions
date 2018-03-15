package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotReader *rot13Reader) Read(b []byte) (n int, e error) {
	n, e = rotReader.r.Read(b)
	for i := range b {
		c := b[i]
		switch {
			case (c >= 'A' && c < 'N') || (c >= 'a' && c < 'n'):
				b[i] = c + 13
			case (c >= 'N' && c <= 'Z') || (c >= 'n' && c <= 'z'):
				b[i] = c - 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}