package main

import (
	"fmt"
	"os"
)

func main() {

	city := os.Args[1]

	switch city {
	case "Colombo":
		fmt.Println("Sri Lanka")
	case "Auckland":
		fmt.Println("New Zealand")
	default:
		fmt.Println("Where ??")
	}
}
