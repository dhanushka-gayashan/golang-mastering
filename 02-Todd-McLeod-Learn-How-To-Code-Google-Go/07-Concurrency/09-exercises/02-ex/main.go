package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("My name is", p.name)
}

func saySomething(h human)  {
	h.speak()
}

func main() {

	p := person{name: "Jack"}
	p.speak()

	// saySomething(p) -> won't work
	saySomething(&p)
}
