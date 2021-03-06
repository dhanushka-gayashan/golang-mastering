package main

import "fmt"

func main() {

	var any interface{}

	// The variable can accept any type of value.
	any = []int{1, 2, 3}
	any = map[int]bool{1: true, 2: false}
	any = "hello"
	any = 3

	// You can't multiply the last number.
	// Reason: `any` is an `interface{}`, it's not a number.
	// int is dynamic value
	any = any.(int) * 2
	fmt.Println(any)
}
