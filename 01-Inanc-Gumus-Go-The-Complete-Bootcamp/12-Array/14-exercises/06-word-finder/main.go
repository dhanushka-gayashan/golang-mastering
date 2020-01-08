package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please give me a word to search.")
		return
	}

	avoid := [...]string{"and", "or", "was", "the", "since", "very"}

	wordList := strings.Fields(strings.ToLower(corpus))

main:
	for _, w := range args {
		for _, aw := range avoid {
			if aw == w {
				continue main
			}
		}

		for i, lw := range wordList {
			if lw == w {
				fmt.Printf("#%-2d : %q\n", i+1, w)
				break

			}
		}
	}
}
