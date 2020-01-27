package main

import "fmt"

func main() {

	i := foo()
	fmt.Println(i)

	xi, xs := bar()
	fmt.Println(xi, xs)
}

func foo() int {
	return 40
}

func bar() (int, string) {
	return 30, "Hello Golang"
}
