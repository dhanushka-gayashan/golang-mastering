package main

import "fmt"

func main() {

	total := sum(1, 2, 3, 4, 5, 6)

	fmt.Println("Total is ", total)
}

func sum(x ...int) int {

	var sum int

	for _, v := range x {
		sum += v
	}

	return sum
}
