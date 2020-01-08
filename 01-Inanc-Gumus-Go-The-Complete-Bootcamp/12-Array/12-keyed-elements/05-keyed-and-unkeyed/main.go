package main

import "fmt"

func main() {

	rates := [...]float64{
		5: 1.5,
		0.5,
		0: 2.5,
	}

	fmt.Println("Length of rates : ", len(rates))
	fmt.Println(rates)
}
