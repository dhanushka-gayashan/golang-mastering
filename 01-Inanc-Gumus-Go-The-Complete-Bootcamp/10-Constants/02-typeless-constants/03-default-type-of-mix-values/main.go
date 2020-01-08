package main

import "fmt"

func main() {

	// Types are compatible, therefore Go will convert both values in to float64
	i := 42 + 2.14
	fmt.Printf("Type of i is %T\n", i)

	// Does not work, becuase data types are incompatible
	//f := true + "Hello"
}
