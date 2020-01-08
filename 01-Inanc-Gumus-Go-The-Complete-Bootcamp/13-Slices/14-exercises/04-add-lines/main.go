package main

import (
	"fmt"
	"strings"
)

func main() {

	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday`)

	fix := make([]string, len(lyric)+3)

	//copy(fix[:8], lyric[:8])
	//fix[8] = "\n"
	//copy(fix[9:19], lyric[8:18])
	//fix[19] = "\n"
	//copy(fix[20:], lyric[18:])
	//fix[len(fix) - 1] = "\n"
	//fmt.Println(fix)

	// yesterday all my troubles seemed so far away -> 8
	// now it looks as though they are here to stay -> 10
	// oh i believe in yesterday -> 5
	cutpoints := []int{8, 10, 5}

	for i, n := 0, 0; n < len(lyric); i++ {

		n += copy(fix[n+i:], lyric[n:n+cutpoints[i]])

		fix[n+i] = "\n"
	}

	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}
