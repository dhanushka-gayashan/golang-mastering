package main

import "fmt"

func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	// if use "range" we need to find a way to close the channel
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}
