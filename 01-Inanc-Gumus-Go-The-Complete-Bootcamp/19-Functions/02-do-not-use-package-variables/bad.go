package main

import "fmt"

func incrN() {
	N++
}

func showN() {
	if N == 0 {
		return // exist from the function
	}

	fmt.Println("N is ", N)
}
