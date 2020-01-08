package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	errorMessage1 = "Provide only the [starting] and [stopping] positions"
	errorMessage2 = "Wrong positions"
)

func main() {

	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	fmt.Printf("%[1]q\n", ships)

	from, to := 0, len(ships)

	switch query := os.Args[1:]; len(query) {
	default:
		fallthrough
	case 0:
		fmt.Println(errorMessage1)
		return
	case 2:
		to, _ = strconv.Atoi(query[1])
		fallthrough
	case 1:
		to, _ = strconv.Atoi(query[0])
	}

	if l := len(ships); from < 0 || from > l || to > l || from > to {
		fmt.Println(errorMessage2)
		return
	}

	fmt.Println(ships[from:to])
}
