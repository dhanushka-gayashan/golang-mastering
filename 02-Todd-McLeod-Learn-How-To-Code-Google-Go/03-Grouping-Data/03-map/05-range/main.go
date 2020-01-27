package main

import "fmt"

func main() {

	phone := map[string]int{
		"Jack" : 123456,
		"Nick" : 67890,
		"Bob" : 54674,
	}

	for k, v := range phone {
		fmt.Println(k, " - ", v)
	}
}
