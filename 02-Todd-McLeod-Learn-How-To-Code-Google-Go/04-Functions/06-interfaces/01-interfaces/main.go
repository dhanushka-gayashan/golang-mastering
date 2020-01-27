package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func run (h human) {
	fmt.Println("Humans can run...")
}

func main() {

	sa1 := secretAgent{
		person: person{"James", "Bond"},
		ltk:    true,
	}
	sa1.speak()
	run(sa1)


	sa2 := secretAgent{
		person: person{"Miss", "MoneyPenny"},
		ltk:    true,
	}
	sa2.speak()
	run(sa2)
}
