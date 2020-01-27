package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	show(even, si...)
	show(odd, si...)
}

func show(f func(si ...int)int, si ...int) {
	total := f(si...)
	fmt.Println(total)
}

func even(si ...int) int {
	var total int
	for _, v := range si {
		if v % 2 != 0 {
			continue
		}
		total += v
	}
	return total
}

func odd(si ...int) int {
	var total int
	for _, v := range si {
		if v % 2 == 0 {
			continue
		}
		total += v
	}
	return total
}