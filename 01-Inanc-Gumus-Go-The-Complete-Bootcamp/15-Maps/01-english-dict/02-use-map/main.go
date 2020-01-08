package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}

	dict := map[string]string{
		"good":    "kötü",
		"great":   "harika",
		"perfect": "mükemmel",
	}

	dict["good"] = "iyi" //can change exist values
	dict["up"] = "yukarı"
	dict["down"] = "aşağı"
	fmt.Printf("Dictionary : %#v\n", dict)

	query := args[0]
	value := dict[query]
	if value == "" {
		fmt.Printf("Query : %q not found\n", query)
		return
	}

	fmt.Printf("Query : %q Value : %q\n", query, value)
}
