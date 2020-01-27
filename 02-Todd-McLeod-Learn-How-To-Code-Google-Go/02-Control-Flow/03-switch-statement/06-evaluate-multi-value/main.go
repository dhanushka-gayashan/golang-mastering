package main

import "fmt"

func main() {

	n := "Bond"

	switch n {
	case "Hello", "Bond", "Shoot":
		fmt.Println("James Bond")
	case "Bye":
		fmt.Println("Bye All")
	default:
		fmt.Println("This is default")
	}
}
