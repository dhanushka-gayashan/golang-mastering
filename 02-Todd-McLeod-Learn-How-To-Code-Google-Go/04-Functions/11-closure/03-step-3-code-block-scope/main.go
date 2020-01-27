package main

import "fmt"

func main() {

	var x int
	fmt.Println(x)

	x++
	fmt.Println(x)

	{
		y := 40
		fmt.Println(y)
	}
	// fmt.Println(y) -> won't work
}
