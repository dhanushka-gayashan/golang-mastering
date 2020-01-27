package main

import "fmt"

func main() {

	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("Defer foo()")
	}()
	fmt.Println("Call foo()")
}

func bar() {
	fmt.Println("Call bar()")
}
