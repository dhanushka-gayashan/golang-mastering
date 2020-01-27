package main

import "fmt"

func main() {

	f1 := func() {
		fmt.Println("Hello Golang World")
	}
	f1()

	f2 := func(name string) {
		fmt.Println("Hello", name)
	}
	f2("Nick")
}
