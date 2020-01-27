package main

import "fmt"

func main() {

	send := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("from sender:", <-send)
		}
		quit<- 0
	}()

	sender(send, quit)
}

func sender(send, quit chan int) {
	var x int

	for {
		select {
		case send<- x:
			x += 2
		case <-quit:
			fmt.Println("Going to exit from sender")
			return
		}
	}
}
