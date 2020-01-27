package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go sender(c)

	receiver(c)

	fmt.Println("You did it baby..")
}

func sender(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receiver(c <-chan int) {
	for v := range c{
		fmt.Println(v)
	}
}
