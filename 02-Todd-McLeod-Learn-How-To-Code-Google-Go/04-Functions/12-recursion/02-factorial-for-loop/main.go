package main

import "fmt"

func main() {

	f := factorial(4)
	fmt.Println(f)
}

func factorial(n int) int {
	v := 1
	for ; n > 0; n-- {
		v *= n
	}
	return v
}