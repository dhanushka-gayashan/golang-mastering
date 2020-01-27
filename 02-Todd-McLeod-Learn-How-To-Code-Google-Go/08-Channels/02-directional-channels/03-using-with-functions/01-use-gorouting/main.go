package main

import "fmt"

func main() {

	c := make(chan int)

	//Both getting block -> Race Condition / DeadLock
	sender(c)
	receiver(c)

	fmt.Println("All Good...")
}

func sender(c chan<- int) {
	c <- 100
}

func receiver(c <-chan int) {
	fmt.Println(<-c)
}
