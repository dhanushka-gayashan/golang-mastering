package main

import (
	"fmt"
	"strconv"
)

func main() {

	// String + Raw String
	fmt.Println(
		`hello` + `, ` + `how` + ` ` + `are` + ` ` + "today?",
	)

	// String + Integer
	eq := "1 + 2 = "
	sum := 1 + 2
	fmt.Println(eq + strconv.Itoa(sum))

	// String + Bool
	eq = strconv.FormatBool(true) + " " + strconv.FormatBool(false)
	fmt.Println(eq)
}
