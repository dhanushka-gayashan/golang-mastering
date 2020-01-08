package main

import "fmt"

func main() {

	// Go convert 20 to int32
	const min int32 = 10
	max := 20 + min
	fmt.Printf("Type of Max is %T\n", max)

	// Put Default value for const value and second value
	// Default value of integer literal is "int"
	const first = 10
	second := 20 + first
	fmt.Printf("Type of Second is %T\n", second)

	//Default Types of untype literals
	i := 100
	fmt.Printf("type of i is %T\n", i)

	f := 12.12
	fmt.Printf("type of f is %T\n", f)

	b := true
	fmt.Printf("type of b is %T\n", b)

	s := "Hello"
	fmt.Printf("type of s is %T\n", s)

	r := 'A'
	fmt.Printf("type of r is %T\n", r)
}
