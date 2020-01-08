package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// len() calculate numbers of bytes of the name
	// Non english character may have 1-4 bytes per character
	name := "İnanç"
	fmt.Printf("%q is %d bytes.\n", name, len(name))

	// Calculate length correctly
	fmt.Printf("%q is %d bytes.\n", name, utf8.RuneCountInString(name))
}
