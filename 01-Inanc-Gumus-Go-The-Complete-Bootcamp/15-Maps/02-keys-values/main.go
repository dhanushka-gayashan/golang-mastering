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

	numbs[100] = ""

	query := 100

	//value, ok := numbs[query]
	//if !ok {
	//	fmt.Printf("%q not found\n", query)
	//	return
	//}
	//
	//fmt.Printf("%q is %q\n", query, value)

	if value, ok := numbs[query]; ok {
		fmt.Printf("%d is %q\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)
}
