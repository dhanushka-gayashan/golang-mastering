package main

import "fmt"

const (
	_ = iota
	one = 2019 - iota
	two = 2019 - iota
	three = 2019 - iota
	four = 2019 - iota
)

func main() {

	fmt.Println(one, two, three, four)
}
