package main

import "fmt"

func main() {

	data := []float64{10, 25, 30, 50}

	correctData := []float64{99, 100}

	copy(data, correctData)

	fmt.Printf("Is it gonna rain? %.f%% chance.\n", (data[0]+data[1]+data[2]+data[3])/float64(len(data)))
}
