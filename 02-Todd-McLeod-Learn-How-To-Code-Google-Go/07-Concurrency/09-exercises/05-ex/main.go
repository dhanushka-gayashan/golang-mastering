package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var (
		count int64
		wg sync.WaitGroup
		gr = 100
	)

	wg.Add(gr)

	fmt.Println("Start Num of Goroutines:", runtime.NumGoroutine())

	for i := 0; i < gr; i++ {
		go func() {
			runtime.Gosched()

			atomic.AddInt64(&count, 1)
			fmt.Println("Count is ", atomic.LoadInt64(&count))
			fmt.Println("Num of Goroutines:", runtime.NumGoroutine())

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End Num of Goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", count)
}

// go run -race main.go