package main

import "fmt"

func main() {

	students := [...][3]float64{
		{9, 4, 7},
		{5, 7, 4},
		{9, 5, 6},
	}

	var sum float64

	for _, marks := range students {
		for _, mark := range marks {
			sum += mark
		}
	}

	fmt.Printf("Total Sum = %.2f\n", sum)

	const N = float64(len(students) * len(students[0]))
	fmt.Printf("Avarage = %.2f\n", sum/N)
}
