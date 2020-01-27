package main

import "fmt"

func main() {

	x := sum()

	fmt.Println("Total is : ", x)
}

func sum(x ...int) int {

	var sum int

	if len(x) == 0  {
		return sum
	}

	for _, i := range x {
		sum += i
	}

	return sum
}
