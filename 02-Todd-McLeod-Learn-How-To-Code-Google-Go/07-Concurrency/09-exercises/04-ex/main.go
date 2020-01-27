package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var (
		count int
		wg sync.WaitGroup
		mtx sync.Mutex
		gr = 100
	)

	wg.Add(gr)

	fmt.Println("Start Num of Goroutines:", runtime.NumGoroutine())

	for i := 0; i < gr; i++ {

		go func() {
			mtx.Lock()
			v := count

			runtime.Gosched()

			v++
			count = v
			fmt.Println("Num of Goroutines:", runtime.NumGoroutine())

			mtx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End Num of Goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", count)
}

// go run -race main.go