package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	v, ok := <-ch
	fmt.Println(v, ok)
}
