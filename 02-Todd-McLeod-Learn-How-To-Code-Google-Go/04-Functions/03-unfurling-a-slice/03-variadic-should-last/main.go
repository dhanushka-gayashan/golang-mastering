package main

import "fmt"

func main() {

	total := sum("James", 10, 20, 30)
	fmt.Println("Total is : ", total)
}

func sum(name string, x ...int) int {

	fmt.Println("Hi ", name)

	var sum int
	for _, i := range x {
		sum += i
	}

	return sum
}
