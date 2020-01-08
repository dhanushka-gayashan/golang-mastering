package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("Tell me a book title")
		return
	}

	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	query := strings.ToLower(arg[0])

	fmt.Println("Search Results:")

	var findBook bool

	for _, book := range books {
		if strings.Contains(strings.ToLower(book), query) {
			fmt.Printf("+ %s\n", book)
			findBook = true
		}
	}

	if !findBook {
		fmt.Printf("We don't have the book: %q\n", query)
	}
}
