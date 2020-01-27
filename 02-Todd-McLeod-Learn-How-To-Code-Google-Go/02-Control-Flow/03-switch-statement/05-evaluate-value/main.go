package main

import "fmt"

func main() {

	switch "Bond" {
	case "Hello":
		fmt.Println("Hello World")
	case "Bond":
		fmt.Println("James Bond")
	case "Bye":
		fmt.Println("Bye All")
	default:
		fmt.Println("This is default")
	}
}
