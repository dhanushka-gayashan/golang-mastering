package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7, 8}

	fmt.Printf("Length is : %d\n", len(x))
	fmt.Printf("Capacity is : %d\n", cap(x))
}
