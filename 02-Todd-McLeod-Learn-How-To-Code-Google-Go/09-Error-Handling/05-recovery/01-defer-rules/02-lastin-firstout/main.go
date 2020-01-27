package main

import "fmt"

func main() {
	// Deferred function calls are executed in Last In First Out order after the surrounding function returns
	foo()
}

func foo() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
