package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("Num of CPU :", runtime.NumCPU())
	fmt.Println("Num of Goroutine :", runtime.NumGoroutine())

	//shared variable among go routine
	// int64 for package of atomic
	var count int64

	gr := 100
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			runtime.Gosched() // Thread Yield

			fmt.Println("Count\t", atomic.LoadInt64(&count))

			wg.Done()
		}()
		fmt.Println("Num of Goroutine inside For-Loop :", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Num of CPU :", runtime.NumCPU())
	fmt.Println("Num of Goroutine :", runtime.NumGoroutine())
	fmt.Println("Count :", count)
}
