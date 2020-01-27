package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("You did it baby..")
}

func send(even, odd,  quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i % 2 == 2 {
			even <- i
		} else {
			odd <- i
		}
	}

	quit <- 100 // not needed this, just for experiment
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("from the even channel:", v)
		case v := <-odd:
			fmt.Println("from the odd channel:", v)
		case  <-quit:
			fmt.Println("from the quit channel")
			return
		}
	}
}




