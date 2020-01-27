package main

import "fmt"

func main() {

	phone := map[string]int{
		"Jack" : 123456,
		"Nick" : 67890,
		"Bob" : 54674,
	}
	fmt.Println(phone)


	delete(phone, "Bob")
	fmt.Println(phone)


	if _, ok := phone["Nick"]; ok {
		delete(phone, "Nick")
	}
	fmt.Println(phone)
}
