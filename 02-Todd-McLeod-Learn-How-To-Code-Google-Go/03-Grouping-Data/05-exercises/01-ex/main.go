package main

import "fmt"

func main() {

	var arr [5]int

	arr[0] = 00
	arr[1] = 10
	arr[2] = 20
	arr[3] = 30
	arr[4] = 40

	for i, v := range arr {
		fmt.Printf("%d index - %d value\n", i, v)
	}

	fmt.Printf("Type of arr is %T\n", arr)
}

