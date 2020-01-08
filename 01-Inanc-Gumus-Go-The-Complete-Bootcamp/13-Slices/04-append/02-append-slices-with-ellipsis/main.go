package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}
	tens := []int{11, 12, 13}

	nums = append(nums, tens...)

	fmt.Println(nums)
}
