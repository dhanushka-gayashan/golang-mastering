package main

import "fmt"

func main() {

	x := []int{0, 10, 20, 30, 40, 50, 40, 50, 60, 70, 80, 90}
	fmt.Println(x)
	fmt.Println("Length", len(x), "Capacity", cap(x))

	fmt.Println()

	x = append(x[:6], x[8:]...)
	fmt.Println(x)
	fmt.Println("Length", len(x), "Capacity", cap(x))

	fmt.Println()

	x = append(x, 200, 300)
	fmt.Println(x)
	fmt.Println("Length", len(x), "Capacity", cap(x))

	fmt.Println()

	x = append(x, 400, 500)
	fmt.Println(x)
	fmt.Println("Length", len(x), "Capacity", cap(x))
}
