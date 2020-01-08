package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//Change String
	// String -> Read Only Byte Slice -> Immutable
	str := "Yūgen ☯ 💀"

	bytes := []byte(str)
	bytes[0] = 'N'
	bytes[1] = 'O'

	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))

	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

	// Print Bytes of the String
	fmt.Println()
	for i, r := range str {
		fmt.Printf("str[%2d] = % -2x = %q\n", i, string(r), r)
	}

	// Print Runes of the String
	fmt.Println()
	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}
}
