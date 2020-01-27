package main

import "fmt"

func main() {

	myAge := 33

	if myAge > 40 {
		fmt.Println("I am too old......")
	} else if myAge < 40 && myAge > 19 {
		fmt.Println("I am young....")
	} else {
		fmt.Println("I am a teenager")
	}
}
