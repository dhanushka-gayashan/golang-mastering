package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("4 + 5 =", mySum(4, 5))
	fmt.Println("6 + 7 =", mySum(6, 7))
}

func mySum(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

func mySumFail(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum + 1
}
