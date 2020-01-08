package main

import "fmt"

func main() {

	var ages []int

	/*
		Length : The length of a slice
		Capacity : The length of the backing array beginning from the first element of the slice
	*/

	fmt.Println(ages)
	fmt.Println("Length :", len(ages))
	fmt.Println("Capacity :", cap(ages))

	ages = append(ages, 10, 20, 30, 40)

	fmt.Println(ages)
	fmt.Println("Length :", len(ages))
	fmt.Println("Capacity :", cap(ages))

	ages = ages[0:0]

	fmt.Println(ages)
	fmt.Println("Length :", len(ages))
	fmt.Println("Capacity :", cap(ages))
}
