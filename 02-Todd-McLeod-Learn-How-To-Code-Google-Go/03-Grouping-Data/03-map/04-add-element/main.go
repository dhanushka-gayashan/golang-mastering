package main

import "fmt"

func main() {

	phone := map[string]int{
		"Jack" : 123456,
		"Nick" : 67890,
	}

	phone["Bob"] = 54674

	fmt.Println(phone)
}
