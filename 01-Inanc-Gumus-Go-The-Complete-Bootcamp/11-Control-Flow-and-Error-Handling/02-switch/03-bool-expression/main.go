package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	num, _ := strconv.Atoi(os.Args[1])

	switch {
	case num > 0:
		fmt.Println("Positive")
	case num < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
