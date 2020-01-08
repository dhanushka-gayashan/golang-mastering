package main

import (
	"bytes"
	"fmt"
)

func main() {

	png, header := []byte{'P', 'N', 'G'}, []byte{}

	header = append(header, png...)

	ok := "not"

	if bytes.Equal(png, header) {
		ok = ""
	}

	fmt.Printf("PNG and HEADER are %sequal", ok)
}
