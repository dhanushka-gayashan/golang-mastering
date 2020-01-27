package main

import "fmt"

func main() {

	c := make(chan int)

	// Main does not wait until Sender and Receiver goroutines are getting completed, because both run in separate goroutines
	// We have not implemented Wait mechanism here like WaitGroup
	go sender(c)
	go receiver(c)

	fmt.Println("All Good...")
}

func sender(c chan<- int) {
	c <- 100
}

func receiver(c <-chan int) {
	fmt.Println(<-c)
}
