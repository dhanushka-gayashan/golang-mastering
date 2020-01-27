package main

import "fmt"

func main() {
	// Deferred functions may read and assign to the returning function's named return values
	i := foo()
	fmt.Println(i)
}

func foo() (i int) {
	defer func() { i++ }()
	return 1
}
