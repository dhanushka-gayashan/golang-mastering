package main

import "fmt"

func main() {

	myAge := 33

	switch myAge{
	case 40:
		fmt.Println("I am too old......")
	case 33:
		fmt.Println("I am young....")
	case 19:
		fmt.Println("I am a teenager")
	}
}
