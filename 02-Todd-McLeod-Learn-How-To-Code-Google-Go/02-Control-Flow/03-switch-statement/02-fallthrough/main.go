package main

import "fmt"

func main() {

	x := 2
	y := 3
	z := 4

	switch {
	case false:
		fmt.Println("This should not print")
	case (x == z):
		fmt.Println("This should not print")
	case (y == y):
		fmt.Println("This should print")
		fallthrough
	case (z == z):
		fmt.Println("This should print ???")
	}
}
