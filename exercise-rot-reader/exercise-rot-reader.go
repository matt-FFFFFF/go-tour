package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(bytes []byte) (int, error) {
	n, err := rot.r.Read(bytes)

	for i := 0; i < n; i++ {
		switch {
		case bytes[i] >= byte('a') && bytes[i] <= byte('m'):
			bytes[i] = bytes[i] + 13
		case bytes[i] >= byte('n') && bytes[i] <= byte('z'):
			bytes[i] = bytes[i] - 13
		case bytes[i] >= byte('A') && bytes[i] <= byte('M'):
			bytes[i] = bytes[i] + 13
		case bytes[i] >= byte('N') && bytes[i] <= byte('Z'):
			bytes[i] = bytes[i] - 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
