package main

import "fmt"

func main() {

	var (
		nums   []int
		oldCap float64
	)

	for len(nums) < 10e6 {
		newCap := float64(cap(nums))

		if newCap == 0 || newCap != oldCap {
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n", len(nums), newCap, newCap/oldCap)
		}

		oldCap = newCap

		nums = append(nums, 1)
	}
}
