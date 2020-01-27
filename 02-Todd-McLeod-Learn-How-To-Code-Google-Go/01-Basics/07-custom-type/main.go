package main

import "fmt"

var i int

type hotdog int
var h hotdog

func main() {

	i = 50
	fmt.Printf("i is %T type and stored value is %v\n", i, i)

	h= 100
	fmt.Printf("h is %T type and stored value is %v\n", h, h)
}
