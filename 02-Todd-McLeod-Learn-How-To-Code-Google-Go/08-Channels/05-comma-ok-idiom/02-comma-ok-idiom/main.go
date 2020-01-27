package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 100
		close(ch)
	}()

	v, ok := <-ch
	fmt.Println(v, ok)

	v, ok = <-ch
	fmt.Println(v, ok)
}
