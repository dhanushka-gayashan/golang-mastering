package main

import "fmt"

func main() {

	c := make(chan int)

	// Sender
	go func(c chan<- int) {
		for i := 0; i < 100; i++ {
			c <- i + 1
		}
		close(c) // otherwise Deadlock
	}(c)

	// Receiver
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("All good...")
}
