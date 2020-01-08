package main

import "fmt"

func main() {

	// Update same Underlying Array, because Capacity is enough
	game1 := []string{"pacman", "mario", "tetris", "doom"}
	subGame1 := game1[0:2]

	fmt.Println("game", game1)
	fmt.Println("subGame", subGame1)
	fmt.Printf("game's - len: %d cap: %d\n", len(game1), cap(game1))
	fmt.Printf("subGame's - len: %d cap: %d\n\n", len(subGame1), cap(subGame1))

	subGame1 = append(subGame1, "dota")

	fmt.Println("game", game1)
	fmt.Println("subGame", subGame1)
	fmt.Printf("game's - len: %d cap: %d\n", len(game1), cap(game1))
	fmt.Printf("subGame's - len: %d cap: %d\n\n", len(subGame1), cap(subGame1))

	subGame1 = append(subGame1, "vgt")

	fmt.Println("game", game1)
	fmt.Println("subGame", subGame1)
	fmt.Printf("game's - len: %d cap: %d\n", len(game1), cap(game1))
	fmt.Printf("subGame's - len: %d cap: %d\n\n", len(subGame1), cap(subGame1))

	fmt.Println()
	fmt.Println()

	//====================================================================================

	// Create new Underlying Array, because Capacity is not enough
	game2 := []string{"pacman", "mario", "tetris", "doom"}
	subGame2 := game1[0:2]

	fmt.Println("game2", game2)
	fmt.Println("subGame2", subGame2)
	fmt.Printf("game2's - len: %d cap: %d\n", len(game2), cap(game2))
	fmt.Printf("subGame2's - len: %d cap: %d\n\n", len(subGame2), cap(subGame2))

	subGame2 = append(subGame2, "dota", "vgt", "fc", "cod", "bf")

	fmt.Println("game2", game2)
	fmt.Println("subGame2", subGame2)
	fmt.Printf("game2's - len: %d cap: %d\n", len(game2), cap(game2))
	fmt.Printf("subGame2's - len: %d cap: %d\n\n", len(subGame2), cap(subGame2))
}
