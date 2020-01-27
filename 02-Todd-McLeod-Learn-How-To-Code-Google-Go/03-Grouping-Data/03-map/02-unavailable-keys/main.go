package main

import "fmt"

func main() {

	phone := map[string]int{
		"Jack" : 123456,
		"Nick" : 67890,
	}

	fmt.Println(phone["Jack"])
	fmt.Println(phone["Nick"])

	fmt.Println(phone["Rose"])
}
