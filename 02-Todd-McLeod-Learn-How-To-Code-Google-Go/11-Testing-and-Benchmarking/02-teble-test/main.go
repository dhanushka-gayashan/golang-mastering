package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("12 + 13 =", mySum(12, 13))
	fmt.Println("22 + 23 =", mySum(22, 23))
	fmt.Println("32 + 33 =", mySum(32, 33))
	fmt.Println("42 + 43 =", mySum(42, 43))
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
