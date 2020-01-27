package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	l float64
}

type circle struct {
	r float64
}

func (s square) area() float64{
	return s.l * s.l
}

func (c circle) area() float64{
	return math.Pi * c.r * c.r
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		l: 100,
	}
	info(s)

	c := circle{r: 70}
	info(c)
}