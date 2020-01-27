package main

import "fmt"

func main() {
	// A deferred function's arguments are evaluated when the defer statement is evaluated.
	foo()
	bar()
}

func foo() {
	i := 0
	//store 0 for i, because i++ executing after this line.
	defer fmt.Println(i)
	i++
	// defer function executing after return
	return
}

func bar() {
	i := 0
	i++
	defer fmt.Println(i)
	return
}
