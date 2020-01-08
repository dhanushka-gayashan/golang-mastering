package main

import "fmt"

func main() {

	// Convert one type to another type
	var apple int
	var orange int64

	apple = int(orange)
	fmt.Println(apple)

	// int to string
	orange = 65
	color := string(orange)
	fmt.Println(color)

	// byte to string
	color = string([]byte{104, 105})
	fmt.Println(color)
}
