package main

import "fmt"

func main() {
	f := get()
	s := f("James Bond")
	fmt.Println(s)
}

func get() func(s string) string {
	return func(name string) string {
		return "Hello, " + name
	}
}
