package main

import "fmt"

func main() {

	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	length := len(items)

	fmt.Println("items", items)

	//Top 3
	top3 := items[:3]
	fmt.Println("top3", top3)

	//Last 4
	last4 := items[length-4:]
	fmt.Println("last4", last4)

	//mid of Last 4
	mid := last4[1:3]
	fmt.Println("mid", mid)

	//slicing return slice
	fmt.Printf("slicing: %T %[1]q\n", items[2:3])

	//indexing return index
	fmt.Printf("indexing: %T %[1]q\n", items[2])
}
