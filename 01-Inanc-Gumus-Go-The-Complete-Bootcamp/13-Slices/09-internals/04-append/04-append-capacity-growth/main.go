package main

import "fmt"

func main() {

	ages, oldCap := []int{1}, 1.

	for len(ages) < 5e5 {
		ages = append(ages, 1)

		newCap := float64(cap(ages))

		if newCap != oldCap {
			fmt.Printf("len:%-10d cap:%-10g growth:%.2f\n", len(ages), newCap, newCap/oldCap)
		}

		oldCap = newCap
	}
}
