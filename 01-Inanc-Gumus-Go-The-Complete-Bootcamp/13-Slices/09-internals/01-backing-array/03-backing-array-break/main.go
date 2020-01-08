package main

import (
	"fmt"
	"sort"
)

func main() {

	grades := []float64{40, 20, 30, 50, 60, 70}

	//var newGrades []float64
	//newGrades = append(newGrades, grades...)
	newGrades := append([]float64(nil), grades...)

	front := newGrades[0:3]

	sort.Float64s(front)

	fmt.Println("grades:", grades)
	fmt.Println("grades:", newGrades)
}
