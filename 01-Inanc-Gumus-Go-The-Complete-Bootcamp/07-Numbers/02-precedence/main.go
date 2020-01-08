package main

import "fmt"

func main() {

	// Conver Celsius to Fahrenheit
	celsius := 35.

	fahrenheit := (9*celsius + 160) / 5
	fmt.Printf("%gC is %gF\n", celsius, fahrenheit)
}
