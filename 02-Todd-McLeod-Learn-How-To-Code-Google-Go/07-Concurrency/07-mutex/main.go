package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Num of CPU :", runtime.NumCPU())
	fmt.Println("Num of Goroutine :", runtime.NumGoroutine())

	var (
		count int
		wg sync.WaitGroup
		mtx sync.Mutex
	)

	gr := 100
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			mtx.Lock()

			v := count

			runtime.Gosched() // Thread Yield

			v++
			count = v

			mtx.Unlock()
			wg.Done()
		}()
		fmt.Println("Num of Goroutine inside For-Loop :", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Num of CPU :", runtime.NumCPU())
	fmt.Println("Num of Goroutine :", runtime.NumGoroutine())
	fmt.Println("Count :", count)
}
