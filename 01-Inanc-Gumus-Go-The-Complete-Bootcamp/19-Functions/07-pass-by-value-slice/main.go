package main

import "fmt"

func main() {

	stats := []int{10, 5}

	// You need to re-assign to stats
	// Slice share the copy of Slice Header
	add(stats, 2)

	fmt.Print(stats)
}

func add(stats []int, n int) {
	stats = append(stats, n)
}
