package main

import "fmt"

func main() {

	var a [5]int
	fmt.Printf("A array : %#v\n", a)
	fmt.Println("Length of A array : ", len(a))

	var s []int
	fmt.Printf("S slice : %#v\n", s)
	fmt.Println("Length of S slice : ", len(s))
}
