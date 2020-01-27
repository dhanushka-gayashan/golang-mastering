package main

import "fmt"

func main() {

	if x := 40; x == 20 {
		fmt.Printf("%v == 20 : %t", x, x == 20)
	} else if x == 40 {
		fmt.Printf("%v == 40 : %t", x, x == 40)
	} else {
		fmt.Printf("Unknown...")
	}
}
