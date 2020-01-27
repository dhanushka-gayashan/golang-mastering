package main

import "fmt"

func main() {

	person := struct {
		first string
		last string
		age int
	}{
		"Mark",
		"Waugh",
		54,
	}

	fmt.Println(person)
}
