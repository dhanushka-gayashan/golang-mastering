package main

import "fmt"

type computer struct {
	brand string
}

func main() {

	var a, b *computer
	fmt.Print(a == b)

	a = &computer{"Apple"}
	b = &computer{"Apple"}
	fmt.Print(" ", a == b, " ", *a == *b)
}
