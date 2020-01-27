package main

import "fmt"

type hotdog int
var x hotdog

var y int

func main() {
	fmt.Printf("%v and type is %T\n", x, x)

	x = 100
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%v and type is %T\n", y, y)

}
