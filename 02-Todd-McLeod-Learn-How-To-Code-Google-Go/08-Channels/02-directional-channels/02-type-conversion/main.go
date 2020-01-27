package main

import "fmt"

func main() {

	cb := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)
	fmt.Printf("%T\n", cb)
	fmt.Printf("%T\n", cs)
	fmt.Printf("%T\n", cr)

	fmt.Println()
	fmt.Println()

	// General to Specific Conversion
	cs = cb
	cr = cb
	fmt.Printf("%T\n", cs)
	fmt.Printf("%T\n", cr)

	fmt.Println()
	fmt.Println()

	//General to Specific Conversion
	fmt.Printf("%T\n", (chan<- int)(cb))
	fmt.Printf("%T\n", (<-chan int)(cb))
}
