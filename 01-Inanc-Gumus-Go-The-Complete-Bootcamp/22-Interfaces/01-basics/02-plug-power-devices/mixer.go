package main

import "fmt"

type mixer struct {}

func (mixer) draw(power int) {
	fmt.Printf("Mixer is drawing %dkW of electrical power.\n", power)
}