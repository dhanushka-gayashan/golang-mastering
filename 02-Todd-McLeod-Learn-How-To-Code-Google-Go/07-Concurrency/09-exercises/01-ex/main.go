package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go show()
	go func() {
		fmt.Println("Hello Golang... from internal func")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Main Goroutine is finished....")
}

func show() {
	fmt.Println("Hello Golang... from external func")
	wg.Done()
}
