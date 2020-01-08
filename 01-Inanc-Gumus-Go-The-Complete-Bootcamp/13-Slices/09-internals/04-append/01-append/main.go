package main

import "fmt"

func main() {

	ages := []int{35, 15}
	fmt.Printf("Pointer : %p Length : %d Capacity : %d Ages : %d\n", &ages, len(ages), cap(ages), ages)

	ages = append(ages, 10)
	fmt.Printf("Pointer : %p Length : %d Capacity : %d Ages : %d\n", &ages, len(ages), cap(ages), ages)
}
