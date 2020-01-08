package main

import "fmt"

func main() {

	const word = "console"

	for _, r := range word {
		fmt.Printf("Rune : %c\t Decimal : %[1]d\t  Hexadecimals : 0x%[1]x\t Binary : 0b%08[1]b\n", r)
	}

	// Runes
	fmt.Printf("With runes       : %s\n", string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

	// Decimals
	fmt.Printf("with decimals    : %s\n", string([]byte{99, 111, 110, 115, 111, 108, 101}))

	// Hexadecimals
	fmt.Printf("with hexadecimals: %s\n", string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))
}
