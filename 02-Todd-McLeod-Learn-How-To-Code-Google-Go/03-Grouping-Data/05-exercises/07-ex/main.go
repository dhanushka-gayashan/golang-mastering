package main

import "fmt"

func main() {

	peoples := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, people := range peoples {
		for _, detail := range people {
			fmt.Print(detail, "\t")
		}
		fmt.Println()
	}
}
