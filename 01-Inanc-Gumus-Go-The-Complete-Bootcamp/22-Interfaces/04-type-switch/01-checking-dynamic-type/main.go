package main

import "fmt"

func main() {

	n := 100
	s := "Hi"
	f := 200.

	check(n)

	check(s)

	check(f)
}

func check(i interface{}) {
	switch i.(type) {
	case int :
		fmt.Println("It is int type")
	case string:
		fmt.Println("i is string type")
	default:
		fmt.Println("i is neither int or string")
	}
}
