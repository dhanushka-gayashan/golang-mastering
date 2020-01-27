package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7, 8}

	for i, v := range x {
		fmt.Printf("%d index - %d value\n", i, v)
	}
}
