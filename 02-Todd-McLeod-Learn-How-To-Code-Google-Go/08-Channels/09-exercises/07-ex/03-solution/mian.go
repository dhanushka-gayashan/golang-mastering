package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	numChans := 10
	numNums := 10

	go sender(numChans, numNums, ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("You did it baby")
}

func sender(numChans, numNums int, ch chan<- string) {
	var wg sync.WaitGroup

	for i := 0; i < numChans; i++ {
		wg.Add(1)
		go func(channel int) {
			for j := 0; j < numNums; j++ {
				value := j
				ch <- fmt.Sprintf("Channel No : %d Value : %d\n", channel, value)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(ch)
}