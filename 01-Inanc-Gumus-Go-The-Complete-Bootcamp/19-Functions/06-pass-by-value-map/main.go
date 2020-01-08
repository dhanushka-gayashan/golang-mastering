package main

import "fmt"

func main() {

	stats := map[int]int{1: 10, 10: 2}

	// No need to re-assign to the stats
	// Map Internal share copy  Map Variable which is point to same Map Header, there for always point to same underlying key-value data structure
	incrAll(stats)

	fmt.Print(stats)
}

func incrAll(stats map[int]int) {
	for k := range stats {
		stats[k]++
	}
}
