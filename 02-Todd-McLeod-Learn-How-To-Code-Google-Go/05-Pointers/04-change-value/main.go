package main

import "fmt"

func main() {

	a := 40

	b := &a

	*b = 50

	fmt.Println(a)
}
