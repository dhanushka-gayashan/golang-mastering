package main

import "fmt"

func main() {

	var add int

	add = add + 1
	fmt.Println(add)

	add++
	fmt.Println(add)

	min := 10

	min = min - 1
	fmt.Println(min)

	min--
	fmt.Println(min)

	var counter int

	counter++
	counter++
	counter++
	counter--
	fmt.Println(counter)
}
