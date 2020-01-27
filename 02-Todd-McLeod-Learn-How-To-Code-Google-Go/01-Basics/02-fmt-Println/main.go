package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello Golang world")
	fmt.Println(n)
	fmt.Println(err)

	//ignore return values
	nn, _ := fmt.Println("Hello Golang world")
	fmt.Println(nn)
}