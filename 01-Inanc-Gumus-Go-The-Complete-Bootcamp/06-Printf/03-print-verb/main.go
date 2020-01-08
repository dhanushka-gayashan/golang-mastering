package main

import "fmt"

func main() {

	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// Print Values
	// This not confirm the Type-Safety
	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions kms\n", distance)
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Does %s has life? %t\n", planet, hasLife)

	// Print Verbs
	// This confim the Type-Safety
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life ? %v\n", planet, hasLife)

	// Argument Indexing
	fmt.Printf(
		"%v is %v away. Think %[2]v kms!! %[1]v OMG",
		planet, distance,
	)

	// Percision
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
}
