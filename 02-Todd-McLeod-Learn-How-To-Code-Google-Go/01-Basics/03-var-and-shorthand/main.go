package main

import "fmt"

var z int
var x = 45

func main() {

	fmt.Println(x)
	foo()

	y := 50
	fmt.Println(y)

	fmt.Println(z)
	z = 60
	fmt.Println(z)
}

func foo() {
	fmt.Println(x)
}
