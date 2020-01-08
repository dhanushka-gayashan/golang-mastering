package main

// File Scope
import "fmt"

// Package Scope
const ok = true

func main() {

	// Block Scope
	var hello = "Hello..!!"

	fmt.Println(ok, hello)
}
