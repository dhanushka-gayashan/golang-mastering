package main

import (
	"fmt"
	"strings"
)

const (
	start = 'A'
	stop  = 'Z'
)

func main() {

	fmt.Printf("%s\n%s\n", "literal", strings.Repeat("=", 30))
	for n := start; n <= stop; n++ {
		fmt.Printf("%c => %[1]d\n", n)
	}
	fmt.Println()

	fmt.Printf("%-10s %-10s\n%s\n", "literal", "decimal", strings.Repeat("=", 30))
	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d\n", n)
	}
	fmt.Println()

	fmt.Printf("%-10s %-10s %-10s\n%s\n", "literal", "decimal", "hexadecimal", strings.Repeat("=", 30))
	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x\n", n)
	}
	fmt.Println()

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n", "literal", "decimal", "hexadecimal", "encoded", strings.Repeat("=", 45))
	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n)) // % -12x unicode bytes separated from space
	}
	fmt.Println()
}
