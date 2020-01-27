package main

import "fmt"

type count int
var c count

func main() {
	fmt.Printf("%v %T\n", c, c)

	c = 42
	fmt.Println(c)
}
