package main

import "fmt"

type kettle struct {}

func (kettle) draw(power int) {
	fmt.Printf("Kettle is drawing %dkW of electrical power.\n", power)
}
