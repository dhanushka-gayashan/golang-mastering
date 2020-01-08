package main

import "fmt"

func main() {

	avengers := struct {
		title, genre string
		rating       float64
		released     bool
	}{
		"avengers: end game", "sci-fi", 8.9, true,
	}

	fmt.Printf("%T\n", avengers)
}
