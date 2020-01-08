package main

import (
	"fmt"
)

func main() {

	ages := []int{35, 15, 25}
	red, green := ages[0:1], ages[1:3]

	fmt.Println("ages", ages)
	fmt.Println("red", red)
	fmt.Println("green", green)

	/*
		Slice == Slice Header
			- 24 bytes

		Slice Header
			- Pointer : Starting point of memory
			- Length  : Actual Length of the Slice
			- Capacity : Capacity of the Slice

		Age
			- Pointer : 100
				- 35 : 100th memory cell
				- 15 : 101th memory cel
				- 25 : 102th memory cel
			- Length : 3
			- Capacity : 3

		Red
			- Pointer : 100
				- 35 : 100th memory cell
			- Length : 1
			- Capacity : 1
		Green
			- Pointer : 101
				- 15 : 101th memory cel
				- 25 : 102th memory cel
			- Length : 2
			- Capacity : 2

		---------------------------------------------
		var empty []int
			- Pointer : 0
			- Length : 0
			- Capacity : 0
		Not use any memory location. There for it is call as "Nil" value
		Does not have a Backing Array
	*/

}
