package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Sender
	go send(even, odd, quit)

	// Receiver
	receiver(even, odd, quit)

	fmt.Println("You did it baby...")
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	// if remove this, main goroutine become deadlock
	// there should be a way to quit from the receiver method, because even and odd goroutines are open
	quit <- 0
}

func receiver(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("from the even channel:", v)
		case v := <-odd:
			fmt.Println("from the odd channel:", v)
		case v := <-quit:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}
