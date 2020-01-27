package main

import "fmt"

type person struct {
	first string
	last string
	flavors []string
}

func main() {

	p1 := person{
		first:   "Mark",
		last:    "Waugh",
		flavors: []string{"chocolate", "vanilla"},
	}

	p2 := person{
		first:   "Steve",
		last:    "Waugh",
		flavors: []string{"Orange", "Mango"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
