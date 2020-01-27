package main

import "fmt"

func main() {

	x := []int{0, 10, 20, 30, 40}
	fmt.Println(x)

	x = append(x, 50, 60, 70)
	fmt.Println(x)

	y := []int{100, 110, 120, 130, 140}
	x = append(x, y...)
	fmt.Println(x)
}
