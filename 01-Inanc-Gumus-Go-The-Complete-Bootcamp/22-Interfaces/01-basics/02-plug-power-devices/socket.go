package main

import (
	"fmt"
	"math/rand"
)

type socket struct {
	power int
}

func (s *socket) plug(device powerdrawer) error {

	n := rand.Intn(50) + 1

	if s.power - n < 0 {
		return fmt.Errorf("socket is out of power of r %dkW", n)
	}

	s.power -= n
	device.draw(n)

	return nil
}
