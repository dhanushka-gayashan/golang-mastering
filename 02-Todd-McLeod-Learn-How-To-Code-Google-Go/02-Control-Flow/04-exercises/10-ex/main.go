package main

import "fmt"

func main() {

	fmt.Printf("true && true\t\t%t\n", true && true)
	fmt.Printf("true && false\t\t%t\n", true && false)
	fmt.Printf("true || true\t\t%t\n", true || true)
	fmt.Printf("true || false\t\t%t\n", true || false)
	fmt.Printf("!true\t\t\t\t%t\n", !true)
}
