package main

import "fmt"

func main() {

	received := Read(0, 3)

	received = append(received, []int{1, 3}...)

	fmt.Println("api.temps     :", All())
	fmt.Println("main.received :", received)
}
