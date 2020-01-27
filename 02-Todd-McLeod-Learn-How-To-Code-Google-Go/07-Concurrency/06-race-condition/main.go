package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Num of CPU :", runtime.NumCPU())
	fmt.Println("Num of Goroutine :", runtime.NumGoroutine())

	//shared variable among go routine
	var count int

	gr := 100
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			v := count

			runtime.Gosched() // Thread Yield

			v++
			count = v
			wg.Done()
		}()
		fmt.Println("Num of Goroutine inside For-Loop :", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Num of CPU :", runtime.NumCPU())
	fmt.Println("Num of Goroutine :", runtime.NumGoroutine())
	fmt.Println("Count :", count)
}
