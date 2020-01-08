package main

import "fmt"

func main() {

	// float64 * int -> does not work
	// follwoing work because constants are typeless
	const min = 1 + 1
	const pi = 3.14 * min
	fmt.Println(min, pi)

	// Go convert constant value into variable type in runtime. like byte(first)
	const first = 100
	var f float64 = first
	var i int = first
	var b byte = first
	var j int32 = first
	var r rune = first
	fmt.Println(f, i, b, j, r)
}
