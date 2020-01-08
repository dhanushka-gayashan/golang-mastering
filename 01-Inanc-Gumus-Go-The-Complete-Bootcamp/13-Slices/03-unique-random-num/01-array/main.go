package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	const max = 5

	var uniques [max]int

main:
	for i := 0; i < max; {
		n := rand.Intn(max + 1)
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue main
			}
		}

		uniques[i] = n
		i++
	}

	fmt.Println("\n\nuniques: ", uniques)
}
