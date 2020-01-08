package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)

	sItems := items[5:8]

	sort.Strings(sItems)

	fmt.Println()
	fmt.Println("Sorted  :", items)
}
