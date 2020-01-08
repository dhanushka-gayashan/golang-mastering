package main

import "fmt"

func main() {

	rates := [...]float64{
		2: 1.5,
		0: 0.5,
		1: 2.5,
	}

	fmt.Printf("%.2f\n", rates[0])
	fmt.Printf("%.2f\n", rates[1])
	fmt.Printf("%.2f\n", rates[2])
}
