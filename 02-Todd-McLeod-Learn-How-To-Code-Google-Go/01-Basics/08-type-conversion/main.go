package main

import "fmt"

var i int = 50

type hotdog int
var h hotdog = 100

func main() {

	fmt.Printf("i is %T type and stored %v\n", i, i)
	fmt.Printf("h is %T type and stored %v\n", h, h)

	//i = h doesn't work. Use Conversion
	i = int(h)
	fmt.Printf("i is %T type and stored %v\n", i, i)
}
