package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	fanin := make(chan int)

	go send(even, odd, quit)

	go receive(even, odd, quit, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("You did it baby")
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(quit)
}

func receive(even, odd, quit <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup

	loop:
		for {
			select {
			case v := <-even:
				wg.Add(1)
				go func() {
					fanin <- v
					wg.Done()
				}()
			case v := <-odd:
				wg.Add(1)
				go func() {
					fanin <- v
					wg.Done()
				}()
			case <- quit:
				break loop
			}
		}
		
	wg.Wait()
	close(fanin)
}