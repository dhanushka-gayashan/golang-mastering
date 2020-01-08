package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var nums []int

	for cap(nums) <= 128 {

		n := rand.Intn(9) + 1
		nums = append(nums, n)

		fmt.Printf("Pointer %p Length %d Capacity %d Numbers %d\n", &nums, len(nums), cap(nums), nums)

		time.Sleep(time.Second / 4)
	}
}
