package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type byAge []person

func (b byAge) Len() int {
	return len(b)
}

func (b byAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byAge) Less(i, j int) bool {
	return b[i].age < b[j].age
}

func main() {

	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)

	sort.Sort(byAge(people))
	fmt.Println(people)
}
