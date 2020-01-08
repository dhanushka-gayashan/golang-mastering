package main

func main() {

	// Purpose of Aliased Type -> Readability
	var (
		byteVal  byte
		uint8Val uint8
	)

	uint8Val = byteVal
	byteVal = uint8Val

	var (
		int32Val int32
		runeVal  rune
	)

	runeVal = int32Val
	int32Val = runeVal

}
