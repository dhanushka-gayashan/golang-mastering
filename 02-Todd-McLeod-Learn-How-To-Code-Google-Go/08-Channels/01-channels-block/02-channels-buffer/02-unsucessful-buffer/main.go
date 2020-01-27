package main

import "fmt"

func main() {

	ch := make(chan int, 1)

	ch <- 100
	ch <- 200 //no room for new value

	fmt.Println(<- ch)
}
