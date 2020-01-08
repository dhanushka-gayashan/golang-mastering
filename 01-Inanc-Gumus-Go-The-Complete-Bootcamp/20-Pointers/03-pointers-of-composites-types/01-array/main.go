package main

import "fmt"

func main() {

	arrays()
}

func arrays() {

	nums := [...]int{1, 2, 3}

	incr(nums)
	fmt.Printf("arrays nums   : %p\n", &nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)
}

func incr(nums [3]int) {
	fmt.Printf("incr nums     : %p\n", &nums)
	for i := range nums {
		nums[i]++
		fmt.Printf("incr.nums[%d]  : %p\n", i, &nums[i])
	}
}

func incrByPtr(nums *[3]int) {
	fmt.Printf("incrByPtr nums: %p\n", &nums)
	for i := range nums {
		nums[i]++ // same: (*nums)[i]++
	}
}
