package main

import "fmt"

func main() {

	n, err := fmt.Println("Hello")

	if err == nil{
		fmt.Println("No any error baby")
		fmt.Printf("%d bytes printed\n", n)
	}
}
