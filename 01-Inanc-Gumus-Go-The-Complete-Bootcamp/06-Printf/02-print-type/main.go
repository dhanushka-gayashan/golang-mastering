package main

import "fmt"

func main() {

	var (
		speed int
		heat  float64
		off   bool
		brand string
	)

	fmt.Printf("Speed is %T\n", speed)
	fmt.Printf("Heat is %T\n", heat)
	fmt.Printf("Off is %T\n", off)
	fmt.Printf("Brand is %T\n", brand)
}
