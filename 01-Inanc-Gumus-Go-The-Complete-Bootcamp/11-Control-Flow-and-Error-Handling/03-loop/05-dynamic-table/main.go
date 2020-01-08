package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const (
		ops   = `+-*/%`
		usage = `Usage: [` + ops + `] [size]`
	)

	switch len(os.Args) {
	case 2:
		fmt.Println("Size is missing")
		fallthrough
	case 1:
		fmt.Println(usage)
		return
	}

	operator := os.Args[1]
	if strings.IndexAny(ops, operator) == -1 {
		fmt.Printf("Invalid operator\n%s", usage)
		return
	}

	rawSize := os.Args[2]
	size, err := strconv.Atoi(rawSize)
	if err != nil {
		fmt.Printf("%q is not a valid Size\n%s", rawSize, usage)
		return
	}

	// First Line
	fmt.Printf("%5s", operator)
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	// Utilize Values
	for i := 0; i <= size; i++ {

		fmt.Printf("%5d", i)

		for j := 0; j <= size; j++ {

			value := 0

			switch operator {
			case "+":
				value = i + j
			case "-":
				value = i - j
			case "*":
				value = i * j
			case "/":
				if j != 0 {
					value = i / j
				}
			case "%":
				if j != 0 {
					value = i % j
				}
			}

			fmt.Printf("%5d", value)
		}

		fmt.Println()
	}
}
