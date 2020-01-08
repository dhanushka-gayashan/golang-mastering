package main

import "fmt"

func main() {

	prev := [3]string{
		"Complete Java Programming",
		"Complete Scala Programming",
		"Complete Golang Programming",
	}

	var books [4]string

	for i, b := range prev {
		books[i] = b + " 2nd Ed."
	}

	books[3] = "Complete Python Programming"

	fmt.Printf("last year: \n%#v\n", prev)
	fmt.Printf("this year: \n%#v\n", books)
}
