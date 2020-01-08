package main

import "fmt"

type blender struct {}

func (blender) draw(power int) {
	fmt.Printf("Blender is drawing %dkW of electrical power.\n", power)
}
