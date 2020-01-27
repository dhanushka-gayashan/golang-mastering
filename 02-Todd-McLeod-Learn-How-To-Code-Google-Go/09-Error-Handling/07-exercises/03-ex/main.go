package main

import "fmt"

func main() {
	cr := customErr{"Custom Error"}
	foo(cr)
	bar(cr)
}

func foo(err error) {
	fmt.Println(err)
}

func bar(err error) {
	fmt.Println(err.(customErr).info)
}
