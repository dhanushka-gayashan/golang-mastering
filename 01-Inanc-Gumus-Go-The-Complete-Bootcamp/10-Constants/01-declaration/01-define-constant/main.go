package main

import "fmt"

func main() {

	// variable should initialize
	// sort-out magic number issue
	const meters int = 100

	cm := 100
	m := cm / meters
	fmt.Printf("%dcm  is %dm\n", cm, m)

	cm = 200
	m = cm / meters
	fmt.Printf("%dcm  is %dm\n", cm, m)
}
