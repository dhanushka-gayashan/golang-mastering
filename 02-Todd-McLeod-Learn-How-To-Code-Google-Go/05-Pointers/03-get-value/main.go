package main

import "fmt"

func main() {

	a := 50

	aa := &a

	va := *aa

	fmt.Println(va)

	fmt.Println(*&a)
}
