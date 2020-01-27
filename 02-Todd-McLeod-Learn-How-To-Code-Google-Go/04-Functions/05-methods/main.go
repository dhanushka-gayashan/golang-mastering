package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func (sg secretAgent) speak() {
	fmt.Println("I am", sg.first, sg.last)
}

func main() {

	sa1:= secretAgent{
		person: person{"James", "Bond"},
		ltk:    true,
	}

	sa2:= secretAgent{
		person: person{"Miss", "MoneyPenny"},
		ltk:    true,
	}

	sa1.speak()
	sa2.speak()
}
