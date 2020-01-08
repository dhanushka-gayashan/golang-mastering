package main

import "fmt"

func main() {

	name := "Google"
	fmt.Printf("%q\n", name)
	fmt.Printf("%s\n", name)

	total, ok, fail := 2500, 2300, 200
	fmt.Printf(
		"total %d - success %d and fail %d",
		total, ok, fail,
	)
}
