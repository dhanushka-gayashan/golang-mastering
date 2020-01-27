package main

import "fmt"

func main() {

	phone := map[string]int{
		"Jack" : 123456,
		"Nick" : 67890,
	}

	if v, ok := phone["Jack"]; ok {
		fmt.Println("Jack", v)
	} else {
		fmt.Println("Jack not available...")
	}
	
	if v, ok := phone["Rose"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Rose not available...")
	}
}
