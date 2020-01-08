package main

import "fmt"

func main() {

	type bookcase [3]int

	blueBooks := bookcase{3, 5, 7}
	redBooks := [3]int{3, 5, 7}

	fmt.Println(blueBooks, " == ", redBooks, " ? ", blueBooks == redBooks)
}
