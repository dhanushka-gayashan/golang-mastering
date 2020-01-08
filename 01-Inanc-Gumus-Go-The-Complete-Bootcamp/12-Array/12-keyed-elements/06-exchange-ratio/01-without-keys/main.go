package main

import "fmt"

func main() {

	rates := [...]float64{
		25.5,
		120.5,
	}

	fmt.Printf("1 BTC is %g ETH\n", rates[0])
	fmt.Printf("1 BTC is %g WAN\n", rates[1])
}
