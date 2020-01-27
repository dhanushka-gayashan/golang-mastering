package main

import "fmt"

func main() {

	if 2 == 2 {
		fmt.Println("Hello, Golang World..")
	}

	if 2 != 2 {
		fmt.Println("Golang is not awesome..")
	}

	if !(2 == 2) {
		fmt.Println("Golang is not awesome..")
	}

	if !(2 != 2) {
		fmt.Println("Golang is hard..")
	}

	if 2 < 1 {
		fmt.Println("Golang is hard..")
	}

	if 2 > 1 {
		fmt.Println("Bye, Golang World..")
	}
}
