package main

import "fmt"

func main() {
	a := currentIndex()
	b := currentIndex()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
}

func currentIndex() func()int {
	var i int
	return func() int {
		i++
		return i
	}
}
