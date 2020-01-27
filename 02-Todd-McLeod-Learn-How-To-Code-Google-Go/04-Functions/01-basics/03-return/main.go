package main

import "fmt"

func main() {

	s := woo("Golang")
	fmt.Println(s)
}

func woo(s string) string {
	return fmt.Sprint("Hello ", s)
}
