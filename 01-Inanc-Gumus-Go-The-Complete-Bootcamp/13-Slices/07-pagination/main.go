package main

import "fmt"

func main() {

	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	const pageSize = 4

	length := len(items)

	for from := 0; from < length; from += pageSize {

		to := from + pageSize
		if to > length {
			to = length
		}

		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d", (from/pageSize)+1)

		fmt.Println(head)
		fmt.Println(currentPage)
	}
}
