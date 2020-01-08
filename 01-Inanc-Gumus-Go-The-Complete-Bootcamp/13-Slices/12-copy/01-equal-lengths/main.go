package main

import "fmt"

func main() {

	nums := []int{1, 2}
	tenth := []int{10, 20}

	copies := copy(nums, tenth)

	fmt.Println("Nums :", nums)
	fmt.Println("Tenth :", tenth)
	fmt.Println("Copied :", copies, "elements")
}
