package main

// It is allow to define unuse variables outside of blocks.
// This is for use this variable in other files or use later.
var outOfBlock int

func main() {

	// It is not allow to unuse.
	var insideBlock int

	// To avoid compile time error, we can use blank identifier
	_ = insideBlock
}
