package main

import "fmt"

const a = 40
const b int = 50

const (
	c = 60
	d int = 70
)

func main() {
	fmt.Printf("%v and %T\n", a, a)
	fmt.Printf("%v and %T\n", b, b)
	fmt.Printf("%v and %T\n", c, c)
	fmt.Printf("%v and %T\n", d, d)
}
