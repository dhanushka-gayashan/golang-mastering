package main

import "fmt"

func main() {

	blueShelf := [...]int{6, 9, 3}

	redShelf := blueShelf
	redShelf[0] = 10
	fmt.Println(blueShelf, redShelf)
}
