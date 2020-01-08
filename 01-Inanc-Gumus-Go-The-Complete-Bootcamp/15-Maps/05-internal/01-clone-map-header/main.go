package main

import "fmt"

func main() {

	english := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
	}

	// Wrong Way
	turkish := english // This copy the Pointer of the English Map. Not the Data

	english["good"] = "do not know"
	turkish["great"] = "out of range"

	// Both point to same memory location in Map Header
	// Both Maps are exist in same memory location
	// Map Variable -> Map Header -> Real Data Set
	fmt.Println(english)
	fmt.Println(turkish)

}
