package main

import "fmt"

func main() {

	x := 1

	for {
		if x > 5 {
			break
		}
		fmt.Println(x)
		x++
	}

	fmt.Println("Done")
}
