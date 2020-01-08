package main

import "fmt"

func main() {

	ages := []int{35, 15, 25}
	first := ages[0:1]
	last2 := ages[1:3]

	ages[0] = 55
	ages[1] = 10
	ages[2] = 20

	grades := []int{70, 99}
	grades[0] = 50

	fmt.Println("ages", ages)
	fmt.Println("first", first)
	fmt.Println("last2", last2)
	fmt.Println("grades", grades)
}
