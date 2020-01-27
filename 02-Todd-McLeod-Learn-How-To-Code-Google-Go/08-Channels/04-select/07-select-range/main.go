package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go sender(even, odd)

	go receiver(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("You did it baby")
}

func sender(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
}

func receiver(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}