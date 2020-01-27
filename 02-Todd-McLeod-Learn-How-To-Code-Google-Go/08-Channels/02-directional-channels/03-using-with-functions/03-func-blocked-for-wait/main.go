package main

import "fmt"

func main() {

	c := make(chan int)

	go sender(c)

	// Execution blocked here. Automatically wait
	receiver(c)

	fmt.Println("All Good...")
}

func sender(c chan<- int) {
	c <- 100
}

func receiver(c <-chan int) {
	//Waiting for receive data from the channel
	fmt.Println(<-c)
}
