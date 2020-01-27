package main

import "fmt"

const (
	a int = 100
	b = "Hello"
)

func main() {

	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
}
