package main

import "fmt"

func main() {

	english := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel",
	}

	fmt.Println(english)

	// Delete 1 Value : Exist and Non-Exist
	delete(english, "awesome")
	delete(english, "awesome")
	delete(english, "ex-awesome")

	fmt.Println(english)

	// Delete all Values
	for k := range english {
		delete(english, k)
	}
}
