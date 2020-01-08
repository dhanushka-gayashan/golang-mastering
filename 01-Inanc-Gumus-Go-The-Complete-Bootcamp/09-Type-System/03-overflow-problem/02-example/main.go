package main

import (
	"fmt"
	"math"
)

func main() {

	// Overflow int8
	var n int8 = math.MaxInt8
	fmt.Println("Max int8: ", n)
	fmt.Println("Max Int8 + 1: ", n+1)

	n = math.MinInt8
	fmt.Println("Min int8: ", n)
	fmt.Println("Min int8 - 1: ", n-1)

	// Overflow uint8
	var un uint8
	fmt.Println("Min uint8: ", un)
	fmt.Println("Min uint8 - 1: ", un-1)

	un = math.MaxUint8
	fmt.Println("Max uint8: ", un)
	fmt.Println("Max uint8 + 1: ", un+1)

	// Overflow float32
	f := float32(math.MaxFloat32)
	fmt.Println("Max float32: ", f*1.2)
}
