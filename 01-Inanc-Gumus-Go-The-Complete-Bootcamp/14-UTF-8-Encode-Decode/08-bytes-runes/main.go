package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const word = "gökyüzü"

	bword := []byte(word)

	fmt.Println(utf8.RuneCount(bword)) // Runes -> 7
	fmt.Println(len(word))             // Bytes -> 10
	fmt.Println(len(string(word[1])))  // Bytes -> 2  (ö -> 2 Bytes | ü -> 2 Bytes)
}
