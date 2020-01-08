package main

import "fmt"

func main() {

	// Byte Max and Min
	var b byte

	b = 0
	fmt.Printf("%08b = %d\n", b, b)

	b = 255
	fmt.Printf("%08b = %d\n", b, b)
}
