package main

import "fmt"

func main() {

	var ratio float64 = 3 / 2
	fmt.Println(ratio)

	ratio = float64(int(3) / int(2))
	fmt.Println(ratio)

	ratio = float64(3) / 2
	fmt.Println(ratio)
}
