package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	switch num, _ := strconv.Atoi(os.Args[1]); /*true is default value going here*/ {
	case num > 0:
		fmt.Println("Positive")
	case num < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
