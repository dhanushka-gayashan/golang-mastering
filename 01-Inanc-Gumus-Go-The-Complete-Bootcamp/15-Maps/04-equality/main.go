package main

import "fmt"

func main() {

	one := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	two := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	first := fmt.Sprintf("%s", one) // "Sprintf" -> Return Formatted String
	second := fmt.Sprintf("%s", two)

	if first == second {
		fmt.Println("Maps are equal...")
	}
}
