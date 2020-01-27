package main

import "fmt"

func main() {

	si := []int{1,2,3,4,5,6,7,8,9}

	total := sum(si...)
	fmt.Println("Total : ", total)

	evenTotal := even(sum, si...)
	fmt.Println("Total Even : ", evenTotal)

	oddTotal := odd(sum, si...)
	fmt.Println("Total Odd", oddTotal)
}

func sum(x ...int) int {

	var total int
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(si ...int) int, si ...int) int {

	var se []int
	for _, v:= range si {
		if v % 2 == 0 {
			se = append(se, v)
		}
	}
	return f(se...)
}

func odd(f func(si ...int) int, si ...int) int {

	var so []int
	for _, v := range si {
		if v % 2 != 0 {
			so = append(so, v)
		}
	}
	return f(so...)
}

