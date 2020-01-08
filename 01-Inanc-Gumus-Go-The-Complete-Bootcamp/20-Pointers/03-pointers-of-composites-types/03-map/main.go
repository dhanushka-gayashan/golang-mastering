package main

import "fmt"

func main() {

	maps()
}

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fix(confused)
	fmt.Println(confused)

	// &confused["one"]
}

// Map values are already a pointer
// No need to pass explicitly
// Unaddressable
func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}
