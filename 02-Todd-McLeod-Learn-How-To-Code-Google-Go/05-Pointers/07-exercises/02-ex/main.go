package main

import "fmt"

type person struct {
	first string
	last string
}

func (p *person)changeMe(s string) {
	p.first = s  // (*p).first
}

func (p person)view() {
	fmt.Printf("First Name : %s Last Name : %s\n", p.first, p.last)
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
	}
	p.view()

	p.changeMe("Jack")
	p.view()
}
