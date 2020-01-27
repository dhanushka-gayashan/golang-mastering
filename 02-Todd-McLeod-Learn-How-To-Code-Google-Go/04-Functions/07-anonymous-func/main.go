package main

import "fmt"

func main() {

	func() {
		fmt.Println("This is Anonymous func....")
	}()

	func(x int) {
		fmt.Println("The meaning of life", x)
	}(42)
}
