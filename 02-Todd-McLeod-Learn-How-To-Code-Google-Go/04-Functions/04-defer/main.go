package main

import "fmt"

func main() {

	// Run as final function
	defer foo()

	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
