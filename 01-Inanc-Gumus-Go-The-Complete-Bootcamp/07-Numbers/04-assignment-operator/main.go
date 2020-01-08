package main

import "fmt"

func main() {

	width, height := 5., 12.

	area := width * height
	fmt.Printf("%g x %g = %g\n", width, height, area)

	// Assignment Operator
	area += 40
	fmt.Printf("%g\n", area)

	area -= 10
	fmt.Printf("%g\n", area)

	area /= 3
	fmt.Printf("%g\n", area)

	area *= 2
	fmt.Printf("%g\n", area)
}
