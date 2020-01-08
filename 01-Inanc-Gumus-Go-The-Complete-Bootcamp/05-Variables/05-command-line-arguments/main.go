package main

import (
	"fmt"
	"os"
)

func main() {

	// Go Build Command accept custom name for generated executable file
	// 	- go build -o greeter
	//  - ls - al
	//  - ./greeter Welcome Golang World
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path to exe file:", os.Args[0])
	fmt.Println("First Argument:", os.Args[1])
	fmt.Println("Second Argument:", os.Args[2])
	fmt.Println("Thired Argument:", os.Args[3])

	fmt.Println("Total Arguments:", len(os.Args))
}
