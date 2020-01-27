package main

import "fmt"

func main() {

	i1 := 10
	i2 := 20

	a := i1 == i2
	b := i1 <= i2
	c := i1 >= i2
	d := i1 != i2
	e := i1 < i2
	f := i1 > i2

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
