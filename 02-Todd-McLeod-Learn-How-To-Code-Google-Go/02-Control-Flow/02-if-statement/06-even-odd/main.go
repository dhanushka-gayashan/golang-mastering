package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {

		if i % 2 == 0 {
			fmt.Printf("%v is an even number\n", i)
		} else {
			fmt.Printf("%v is an odd number\n", i)
		}
	}
}
