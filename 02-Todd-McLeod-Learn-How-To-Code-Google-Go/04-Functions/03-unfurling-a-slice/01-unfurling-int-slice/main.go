package main

import "fmt"

func main() {

	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}

	x := sum(xi...)

	fmt.Println("Total is : ", x)
}

func sum(x ...int) int {
	var sum int

	for _, i := range x {
		sum += i
	}

	return sum
}
