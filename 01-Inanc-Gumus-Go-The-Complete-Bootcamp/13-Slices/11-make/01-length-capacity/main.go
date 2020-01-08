package main

import "fmt"

func main() {

	// Prevent expensive operation of creating new backing array when slice is growing that capacity
	nums1 := make([]int, 3)
	fmt.Printf("Length %d Capacity %d Slice %d\n\n", len(nums1), cap(nums1), nums1)

	nums2 := make([]int, 3, 5)
	fmt.Printf("Length %d Capacity %d Slice %d\n\n", len(nums2), cap(nums2), nums2)

	nums3 := make([]int, 0, 5)
	fmt.Printf("Length %d Capacity %d Slice %d\n\n", len(nums3), cap(nums3), nums3)

	nums3 = append(nums3, 100)
	fmt.Printf("Length %d Capacity %d Slice %d\n\n", len(nums3), cap(nums3), nums3)
}
