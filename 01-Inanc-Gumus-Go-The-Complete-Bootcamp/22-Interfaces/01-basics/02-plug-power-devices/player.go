package main

import "fmt"

type player struct {}

func (player) draw(power int)  {
	fmt.Printf("Player is drawing %dkW of electrical power.\n", power)
}
