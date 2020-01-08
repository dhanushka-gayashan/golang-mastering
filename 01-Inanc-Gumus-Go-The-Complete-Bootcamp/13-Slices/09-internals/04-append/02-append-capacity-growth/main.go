package main

import "fmt"

func main() {

	var ages []int
	fmt.Printf("Pointer : %p Length : %d Capacity : %d Ages : %d\n", &ages, len(ages), cap(ages), ages)

	ages = []int{35, 15}
	fmt.Printf("Pointer : %p Length : %d Capacity : %d Ages : %d\n", &ages, len(ages), cap(ages), ages)

	ages = append(ages, 10)
	fmt.Printf("Pointer : %p Length : %d Capacity : %d Ages : %d\n", &ages, len(ages), cap(ages), ages)

	ages = append(ages, 60)
	fmt.Printf("Pointer : %p Length : %d Capacity : %d Ages : %d\n", &ages, len(ages), cap(ages), ages)

	//Duplicate
	ages = append(ages, ages[2:]...)
	fmt.Printf("Pointer : %p Length : %d Capacity : %d Ages : %d\n", &ages, len(ages), cap(ages), ages)

	// Override Slice
	ages = append(ages[:2], 7, 9)
	fmt.Printf("Pointer : %p Length : %d Capacity : %d Ages : %d\n", &ages, len(ages), cap(ages), ages)

	ages = ages[:cap(ages)]
	fmt.Printf("Pointer : %p Length : %d Capacity : %d Ages : %d\n", &ages, len(ages), cap(ages), ages)

}
