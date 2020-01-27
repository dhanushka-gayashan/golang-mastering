package main

import "fmt"

func main() {

	si := []int{10, 20, 30, 40 , 50, 60, 70, 80, 90}

	tf := foo(si...)
	fmt.Println(tf)

	tb := bar(si)
	fmt.Println(tb)
}

func foo(si ...int) int{
	var total int
	for _, v := range si {
		total += v
	}
	return total
}

func bar(si []int) int {
	var total int
	for _, v := range si {
		total += v
	}
	return total
}