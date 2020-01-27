package main

import "fmt"

func main() {

	arr := []int{10, 20, 30, 40 ,50}

	for i, v := range arr {
		fmt.Printf("%d index - %d value\n", i, v)
	}

	fmt.Printf("Type of arr is %T\n", arr)
}

