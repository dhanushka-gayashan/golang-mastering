package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("OS\t\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
}
