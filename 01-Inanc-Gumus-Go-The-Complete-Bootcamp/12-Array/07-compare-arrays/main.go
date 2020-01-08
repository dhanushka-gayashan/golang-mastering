package main

import "fmt"

func main() {

	blueShelf := [3]int{6, 9, 3}

	redShelf := [3]int{6, 9, 3}

	greenShelf := [3]int{3, 9, 6}

	fmt.Println(blueShelf, " == ", redShelf, " ? ", blueShelf == redShelf)
	fmt.Println(blueShelf, " == ", greenShelf, " ? ", blueShelf == greenShelf)

	// Not comparable
	// [3]int{6, 9, 3} with [2]int{6, 9}
	// [3]int{6, 9, 3} with [3]string{"6", "9", "3"}
}
