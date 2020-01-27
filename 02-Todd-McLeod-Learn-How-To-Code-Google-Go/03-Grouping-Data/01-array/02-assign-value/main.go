package main

import "fmt"

func main() {

	var x [5]int
	fmt.Println(x)

	x[2]= 20
	fmt.Println(x)

	x[4]= 40
	fmt.Println(x)

	x[3]= 30
	fmt.Println(x)

	x[1]= 10
	fmt.Println(x)

	x[0]= 100
	fmt.Println(x)
}
