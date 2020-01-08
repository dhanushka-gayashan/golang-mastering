package main

var temps = []int{5, 10, 3, 25, 45, 80, 90}

func Read(start, stop int) []int {
	// ----------------------------------------
	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA

	portion := temps[start:stop:stop]

	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA
	// ----------------------------------------

	return portion
}

func All() []int {
	return temps
}
