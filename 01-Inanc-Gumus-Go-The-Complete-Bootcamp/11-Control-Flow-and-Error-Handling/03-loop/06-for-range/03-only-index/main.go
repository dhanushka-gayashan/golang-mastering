package main

import (
	"fmt"
	"strings"
)

func main() {

	words := strings.Fields("lazy dog jumps over the gate")

	for i := range words {
		fmt.Println(i)
	}
}
