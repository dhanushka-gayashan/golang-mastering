package main

import "fmt"

func main() {

	numbs := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	// Printing order will be different, because Map optimized for Fast-Looksup not for Fast-Ordered-Looping
	for k, v := range numbs {
		fmt.Printf("%q means %#v\n", k, v)
	}
}
