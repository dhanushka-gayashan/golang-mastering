package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, w := range words {
		// Bytes Length and Runes Length of Strings
		fmt.Printf("Length of Bytes : %d\t Length of Runes :%d\n", len(w), utf8.RuneCountInString(w))

		// Bytes in Hexadecimal
		fmt.Printf("% x\n", w) // []bytes(s) == s

		// Runes in Hexadecimal
		for i, r := range w {
			if i == 0 {
				fmt.Printf("%x", r)
				continue
			}
			fmt.Printf("% x", r)
		}
		fmt.Println()

		// Print Rune
		for _, r := range w {
			fmt.Printf("%q", r) // %c
		}
		fmt.Println()

		//First Rune and Byte Size of the Rune
		firstRune, size := utf8.DecodeRuneInString(w)
		fmt.Printf("First Rune is : %q and Byte Size is : %d\n", firstRune, size)

		// Last Rune
		lastRune, size := utf8.DecodeLastRuneInString(w)
		fmt.Printf("Last Rune is : %q and Byte Size is : %d\n", lastRune, size)

		// First 2 Runes
		_, first := utf8.DecodeRuneInString(w)
		_, second := utf8.DecodeRuneInString(w[first:])
		fmt.Printf("First 2 Runes : %q\n", w[:first+second])

		// Last 2 Runes
		_, last1 := utf8.DecodeLastRuneInString(w)
		_, last2 := utf8.DecodeLastRuneInString(w[:len(w)-last1])
		fmt.Printf("Last 2 Runes : %q\n", w[(len(w)-(last1+last2)):])

		// Convert the string to []rune
		// Print First and Last 2 Runes
		r := []rune(w)
		fmt.Printf("First 2 Rune : %q  Last 2 Rune : %q\n", string(r[:2]), string(r[len(r)-2:]))

		fmt.Println(strings.Repeat("=", 50))
		fmt.Println()
	}
}
