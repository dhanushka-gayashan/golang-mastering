package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	first string
	ltk bool
}

func main() {

	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   40,
		},
		first: "007",
		ltk:    true,
	}

	fmt.Println(sa.first)
	fmt.Println(sa.person.first)
	fmt.Println(sa.last)
	fmt.Println(sa.age)
	fmt.Println(sa.ltk)
}
