package main

import "fmt"

func main() {

	myAge := 33

	switch {
	case myAge > 40:
		fmt.Println("I am too old......")
	case myAge < 40 && myAge > 19:
		fmt.Println("I am young....")
	case myAge < 19:
		fmt.Println("I am a teenager")
	}
}
