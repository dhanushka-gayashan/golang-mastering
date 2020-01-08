package main

import "fmt"

func main() {

	// 'h', 'e', 'y' -> Rune Literal
	// 104, 101, 121 -> Ascii Code Point -> Rune
	// Rune 1 can be 1 or more bytes

	// string to []bytes or []bytes to string
	str := "hey"
	bytes := []byte{104, 101, 121}

	fmt.Printf(`"hey" as bytes   : %d`+"\n", []byte(str))
	fmt.Printf("bytes as string  : %q\n", string(bytes))

	fmt.Println()

	// Rune Literal are Typeless
	var (
		anInt   int   = 'h'
		anInt8  int8  = 'h'
		anInt16 int16 = 'h'
		anInt32 int32 = 'h'

		// rune literal's default type is: rune
		// so, you don't need to specify it like aRune   rune  = 'h'
		aRune = 'h'
	)
	fmt.Printf("rune literals are typeless:\n\t%T %T %T %T %T\n", anInt, anInt8, anInt16, anInt32, aRune)

	fmt.Println()

	// all are the same rune
	// beginning with go 1.13 you can type: 0b0110_1000 instead
	// fmt.Printf("%q as binary: %08[1]b\n", 0b0110_1000)
	fmt.Printf("%q in decimal: %[1]d\n", 104)
	fmt.Printf("%q in binary : %08[1]b\n", 'h')
	fmt.Printf("%q in hex    : 0x%[1]x\n", 0x68)
}
