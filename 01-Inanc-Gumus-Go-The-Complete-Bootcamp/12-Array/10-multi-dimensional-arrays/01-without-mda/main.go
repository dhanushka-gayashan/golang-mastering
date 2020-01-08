package main

import "fmt"

func main() {

	student1 := [3]float64{9, 4, 7}
	student2 := [3]float64{5, 7, 4}
	student3 := [3]float64{9, 5, 6}

	var sum float64
	sum += student1[0] + student1[1] + student1[2]
	sum += student2[0] + student2[1] + student2[2]
	sum += student3[0] + student3[1] + student3[2]

	fmt.Printf("Total Sum = %.2f\n", sum)

	const N = float64(len(student1) * 3)
	fmt.Printf("Avarage = %.2f\n", sum/N)
}
