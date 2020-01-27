package main

import "fmt"

func main() {

	foo := make(chan int)
	bar := make(chan int)

	go func() {
		for i := 0; i < 50; i++ {
			foo <- i
		}
	}()

	go func() {
		for i := 0; i < 50; i++ {
			bar <- i
		}
	}()

	receive(foo, bar)

	fmt.Println("You did it baby")
}

func receive(foo, bar chan int) {
	for i := 0; i < 100; i++ {
		select {
		case v := <-foo:
			fmt.Println("from foo:", v)
		case v := <-bar:
			fmt.Println("from bar:", v)
		}
	}
	fmt.Println("Exit from Receiver()")
}
