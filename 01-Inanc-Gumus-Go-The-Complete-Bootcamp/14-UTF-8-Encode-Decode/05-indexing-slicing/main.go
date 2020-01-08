package main

import (
	"fmt"
	"unsafe"
)

func main() {

	str := "Yūgen ☯ 💀"

	fmt.Printf("1st byte   : %c\n", str[0])           // ok
	fmt.Printf("2nd byte   : %c\n", str[1])           // not ok
	fmt.Printf("2nd rune   : %s\n", str[1:3])         // ok
	fmt.Printf("last rune  : %s\n", str[len(str)-4:]) //ok

	fmt.Println()

	// Inefficient Way
	// String -> 15 bytes
	// Rune -> 9 rune -> 9 * 4 -> 36 bytes (Rune == int32 == 4 bytes)
	runes := []rune(str)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", int(unsafe.Sizeof(runes[0]))*len(runes))
	fmt.Printf("\t%d runes\n", len(runes))

	fmt.Printf("1st rune   : %c\n", runes[0])
	fmt.Printf("2nd rune   : %c\n", runes[1])
	fmt.Printf("first five : %c\n", runes[:5])

	fmt.Println()

	word := "öykü"
	fmt.Printf("%q in runes: %c\n", word, []rune(word))
	fmt.Printf("%q in bytes: % [1]x\n", word)

	fmt.Printf("%s %s\n", word[:2], []byte{word[0], word[1]}) // ö
	fmt.Printf("%c\n", word[2])                               // y
	fmt.Printf("%c\n", word[3])                               // k
	fmt.Printf("%s %s\n", word[4:], []byte{word[4], word[5]}) // ü
}
