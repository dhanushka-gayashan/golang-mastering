package main

import (
	"fmt"
	"time"
)

func main() {

	var (
		pizza      []string
		departure  []time.Time
		graduation []int
		lights     []bool
	)

	now := time.Now()

	pizza = append(pizza, "pepperoni", "onions", "extra", "cheese")
	departure = append(departure, now.Add(time.Hour*24), now.Add(time.Hour*48), now.Add(time.Hour*96))
	graduation = append(graduation, 1998, 2005, 2018)
	lights = append(lights, true, false, true)

	fmt.Printf("Pizza %s\n", pizza)
	fmt.Printf("Departure %s\n", departure)
	fmt.Printf("Graduation %d\n", graduation)
	fmt.Printf("Lights %t\n", lights)
}
