package main

import "fmt"

func main() {

	var (
		width  uint8 = 255
		height       = 255
	)

	width++

	if int(width) < height {
		fmt.Println("Height is greater than Width")
	}

	// Memory of Width has been overflow.
	// uint8 max is 255
	fmt.Printf("Width = %d and Height = %d\n", width, height)
}
