package main

import "fmt"

func main() {

	s, a := person("Golang")
	fmt.Println("Greeting : ", s)
	fmt.Println("Age : ", a)
}

func person(name string) (string, int) {
	s := fmt.Sprintf("Hello %s", name)
	a := 20
	return s, a
}
