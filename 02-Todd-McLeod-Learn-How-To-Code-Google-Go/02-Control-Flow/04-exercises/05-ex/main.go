package main

import "fmt"

func main() {

	for i := 10; i < 100; i++ {
		if m := i % 4; m != 0 {
			fmt.Printf("%d -> %d\n", i, m)
		}
	}
}
