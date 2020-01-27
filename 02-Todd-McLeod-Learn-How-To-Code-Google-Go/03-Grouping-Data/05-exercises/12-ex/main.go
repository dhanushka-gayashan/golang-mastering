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

	ps := map[string]person {
		p1.first: p1,
		p2.first: p2,
	}

	fmt.Println(ps[p1.first])
	fmt.Println(ps[p2.first])
}
