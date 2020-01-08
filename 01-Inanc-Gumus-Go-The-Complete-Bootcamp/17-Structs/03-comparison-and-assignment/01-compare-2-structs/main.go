package main

import "fmt"

type song struct {
	title, artist string
}

func main() {

	song1 := song{title: "wonderwall", artist: "oasis"}
	song2 := song{title: "super sonic", artist: "oasis"}

	fmt.Printf("song1: %+v\nsong2: %+v\n", song1, song2)

	// Struct values are "Bare Values". Go automatically compare all fields in the Structs
	if song1 == song2 {
		fmt.Println("songs are equal.")
	} else {
		fmt.Println("songs are not equal.")
	}
}
