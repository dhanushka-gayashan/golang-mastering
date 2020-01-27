package main

import "fmt"

type animal interface {
	move()
}

type human struct {
	first string
	last string
}

type dog struct {
	name string
	age int
}

func (h human) move() {
	fmt.Println("Human can slowly walk and run")
}

func (h human) laugh() {
	fmt.Println("Human can laugh loudly")
}

func (d dog) move() {
	fmt.Println("Dog can chase a thief...")
}

func (d dog) bark() {
	fmt.Println("Dog can bark loudly")
}

func do(a animal) {

	switch a.(type) {
	case human:
		a.(human).laugh()
	case dog:
		a.(dog).bark()
	}
}

func main() {

	h := human{
		first: "Perter",
		last:  "Jackson",
	}

	d := dog{
		name: "Jimmy",
		age:  5,
	}

	do(h)
	do(d)
}
