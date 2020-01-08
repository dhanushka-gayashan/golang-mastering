package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	usd, err := strconv.ParseFloat(arg[0], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	const (
		EUR = iota
		GBP
		JPY
	)

	rates := [3]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", usd, rates[EUR]*usd)
	fmt.Printf("%.2f USD is %.2f GBP\n", usd, rates[GBP]*usd)
	fmt.Printf("%.2f USD is %.2f JPY\n", usd, rates[JPY]*usd)
}
