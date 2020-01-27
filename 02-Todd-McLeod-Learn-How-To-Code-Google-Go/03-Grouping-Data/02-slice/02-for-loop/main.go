package main

import "fmt"

func main() {

	x := []int{0, 10, 20, 30, 40}

	for i := 0; i < len(x); i++  {
		fmt.Printf("%d index and %d value\n", i, x[i])
	}
}
