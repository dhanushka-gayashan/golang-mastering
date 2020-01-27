package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("You did it baby..")
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i % 2 == 2 {
			even <- i
		} else {
			odd <- i
		}
	}

	quit <- true // not needed this, just for experiment
	close(quit)
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("from the even channel:", v)
		case v := <-odd:
			fmt.Println("from the odd channel:", v)
		case i, ok := <-quit:
			if ok {
				fmt.Println("from comma ok:", i)
			} else {
				fmt.Println("from comma ok:", i, ok)
				return
			}
		}
	}
}




