package main

import (
	"fmt"
	"runtime"
)

func main() {

	// Statement : fmt.Pringln -> Not return a value
	// Expression : runtime.NumCPU() -> Retrun a value
	fmt.Println("No of CPUs : ", runtime.NumCPU())
}
