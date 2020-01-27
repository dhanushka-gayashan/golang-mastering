package main

import "fmt"

type human interface {
	speak()
	run()
}

type person struct {
	first string
	last string
}

func (p person) speak() {
	fmt.Println("Hi,", p.first, p.last, "is a person. He can Speak")
}

func (p person) run() {
	fmt.Println("Hi,", p.first, p.last, "is a person. He can Run")
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("Hi,", sa.first, sa.last, "is a Secret Agent. He can Speak")
}

func main() {

	sa1 := secretAgent{
		person: person{"James", "Bond"},
		ltk:    true,
	}
	sa1.speak()
	sa1.run()


	sa2 := secretAgent{
		person: person{"Miss", "MoneyPenny"},
		ltk:    true,
	}
	sa2.speak()
	sa1.run()
}
