package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   40,
		},
		ltk:    true,
	}

	fmt.Println(sa.first)
	fmt.Println(sa.last)
	fmt.Println(sa.age)
	fmt.Println(sa.ltk)
}
