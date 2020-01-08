package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {

	// Sometimes we need to do decode manual
	word := []byte("öykü")

	fmt.Printf("%s = % [1]x\n", word)

	var size int

	/*
		Option 01:
			- for i := range word -> give bytes by bytes -> But Rune/Letter can be more than 1 byte
			- string(word) -> return Rune by Rune
	*/
	for i := range string(word) {
		if i > 0 {
			size = i
			break
		}
	}

	/*
		Option 02:
			- utf8.DecodeRune(word) -> return first Rune of the word
			- most efficient than for loop
	*/
	_, size = utf8.DecodeRune(word)

	// Change the first Rune(Letter) into Upper Case
	copy(word[:size], bytes.ToUpper(word[:size]))

	fmt.Printf("%s = % [1]x\n", word)
}
