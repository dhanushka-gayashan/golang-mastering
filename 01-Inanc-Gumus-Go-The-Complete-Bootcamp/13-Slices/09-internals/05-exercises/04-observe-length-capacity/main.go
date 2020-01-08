package main

import "fmt"

func main() {

	var game1 []string
	fmt.Printf("Pointer %p Length %d Capacity %d Game %s\n", &game1, len(game1), cap(game1), game1)

	game2 := []string{}
	fmt.Printf("Pointer %p Length %d Capacity %d Game %s\n", &game2, len(game2), cap(game2), game2)

	game2 = append(game2, "pacman", "mario", "tetris", "doom")
	fmt.Printf("Pointer %p Length %d Capacity %d Game %s\n", &game2, len(game2), cap(game2), game2)

	game3 := []string{"pacman", "mario", "tetris", "doom"}
	fmt.Printf("Pointer %p Length %d Capacity %d Game %s\n", &game3, len(game3), cap(game3), game3)

	//----------------------------------------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	//----------------------------------------------------------------------------------------------------

	game4 := []string{"pacman", "mario", "tetris", "doom"}
	for i := 0; i <= len(game4); i++ {
		s := game4[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(s), cap(s))
	}

	//----------------------------------------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	//----------------------------------------------------------------------------------------------------

	test := game4[:2]
	zero := game4[:0]
	fmt.Printf("games's     len: %d cap: %d\n", len(game4), cap(game4))
	fmt.Printf("test's      len: %d cap: %d\n", len(test), cap(test))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	fmt.Println()

	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	//----------------------------------------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	//----------------------------------------------------------------------------------------------------

	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}

	//----------------------------------------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	//----------------------------------------------------------------------------------------------------

	zero = zero[:cap(zero)]
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}

	//----------------------------------------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	//----------------------------------------------------------------------------------------------------

	zero[0] = "command & conquer"
	game4[0] = "red alert"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", game4)
}
