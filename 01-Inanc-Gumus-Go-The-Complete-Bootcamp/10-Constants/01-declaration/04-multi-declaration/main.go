package main

import "fmt"

func main() {

	const min, max int = 10, 20
	fmt.Println(min, max)

	const (
		alpha int = 1
		beta  int = 2
	)
	fmt.Println(alpha, beta)

	// Replace previous type and value
	const (
		apple int = 10
		orange
		iphone int = 10 + 10
		android
	)
	fmt.Println(apple, orange, iphone, android)
}
