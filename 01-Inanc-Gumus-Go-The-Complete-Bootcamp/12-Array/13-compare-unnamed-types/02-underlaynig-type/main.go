package main

import "fmt"

func main() {

	type (
		bookcase [3]int
		cabinet  [3]int
	)

	bc := bookcase{3, 5, 7}
	cb := cabinet{3, 5, 7}

	//fmt.Println(bc, " == ", cb, " ? ", bc == cb) --> does not work, because different type names

	fmt.Println(cabinet(bc), " == ", cb, " ? ", cabinet(bc) == cb)
}
