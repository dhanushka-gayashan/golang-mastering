package main

import "fmt"

func main() {

	nums := []int{1, 3, 2, 4}

	// numbs and odds use same underlying array
	// nums : length 4  capacity 4
	// odds : length 2  capacity 4
	odds := nums[:2]
	odds[0] = 5
	odds[1] = 7

	fmt.Println("nums :", nums)
	fmt.Println("odds :", odds)

	// we can append 2 more to odds slice. Still use same underlying array
	odds = append(odds, 7, 9)

	fmt.Println("nums :", nums)
	fmt.Println("odds :", odds)

	//---------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	//---------------------------------------------------------------------

	nums = []int{1, 3, 2, 4}

	// Limit the Capacity of even slice as 2
	// When add more elements than capacity Go create new underlying array for even slice
	even := append(nums[:2:2], 5, 7)

	fmt.Println("nums :", nums)
	fmt.Println("even :", even)

	//---------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	//---------------------------------------------------------------------

	nums = []int{1, 3, 2, 4}

	//newnums slice references last 2 elements from nums, therefor slice is full. (Length 2 and Capacity 2)
	// When add more elements than capacity Go create new underlying array for even slice
	// This way can use as substitution for full slice expression
	newnums := append(nums[2:4], 5, 7)

	fmt.Println("nums :", nums)
	fmt.Println("newnums :", newnums)
}
