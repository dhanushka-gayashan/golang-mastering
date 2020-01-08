package main

import "fmt"

func main() {

	var safe bool
	weight := 100

	// There should be at least 1 new variable in shorthand declaration
	// Otherwise won't work
	safe, speed := true, 50
	fmt.Println(safe, speed)

	weight, height := 200, 175
	fmt.Println(weight, height)
}
