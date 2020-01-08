package main

import "fmt"

func main() {

	data := []float64{10, 25, 30, 50}

	correctData := []float64{99, 100, 20, 40, 70, 88, 5}

	N := copy(data, correctData) // only copy elements upto the length

	fmt.Printf("%d probabilities copied.\n", N)
	fmt.Println("data :", data)
	fmt.Println("correctData :", correctData)

	fmt.Printf("Is it gonna rain? %.f%% chance.\n", (data[0]+data[1]+data[2]+data[3])/float64(len(data)))
}
