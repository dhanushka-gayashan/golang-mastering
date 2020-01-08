package main

import (
	"fmt"

	"github.com/dhanukdg/Golang-Udemy-Inanc-Gumus-Go-The-Complete-Bootcamp/09-type-system/04-defined-types/02-own-type/weights"
)

type (
	// Gram is a new name type
	Gram int

	// Kilogram is a new name type
	Kilogram Gram

	// Ton is a new name type
	Ton Kilogram
)

func main() {

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)

	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Kilogram(Gram(int(apples))))
	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)

	salt = Gram(weights.Gram(100))

	fmt.Printf("Type of weights.Gram: %T\n", weights.Gram(1))
	fmt.Printf("type of main.Gram: %T\n", Gram(1))

	// New Type from exported Type
	type myGram weights.Gram

	var pepper myGram = 20
	fmt.Printf("Type of pepper: %T\n", pepper)

	// This is correct because underlaying-type is int
	pepper = myGram(salt)
	fmt.Printf("Type of pepper: %T\n", pepper)
}
