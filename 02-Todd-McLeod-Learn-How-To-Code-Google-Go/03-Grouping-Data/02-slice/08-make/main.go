package main

import "fmt"

func main() {

	x := make([]int, 5, 5)
	fmt.Println(x)


	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5
	fmt.Println(x)


	y := make([]int, 5, 5)
	y = append(y, 100, 200, 300)
	fmt.Println(y)
}
