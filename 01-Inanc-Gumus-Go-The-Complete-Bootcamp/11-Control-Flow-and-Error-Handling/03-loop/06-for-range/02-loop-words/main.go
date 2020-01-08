package main

import (
	"fmt"
	"strings"
)

func main() {

	words := strings.Fields("lazy dog jumps over the gate")

	for i, v := range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}
}
