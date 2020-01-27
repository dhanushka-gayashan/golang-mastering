package main

import "fmt"

func main() {

	// Send and Receive should be happen in same time
	ch := make(chan int)

	ch <- 100

	fmt.Println(<-ch)
}