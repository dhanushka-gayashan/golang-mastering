package main

import "fmt"

var myAge = 30

func nestedFunc() {

	var myAge = 40
	fmt.Println("Inside nestedFunc", myAge)
}

func main() {

	fmt.Println("Inside main", myAge)

	nestedFunc()

	fmt.Println("Inside main", myAge)
}
