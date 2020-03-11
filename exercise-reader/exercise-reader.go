package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(bytes []byte) (int, error) {
	//bytes = bytes[:cap(bytes)]
	for i := range bytes {
		bytes[i] = 'A'
	}
	return len(bytes), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
