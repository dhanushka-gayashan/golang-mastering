package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {

	var books [yearly]string

	books[0] = "Complete Java"
	books[1] = "Complete Scala"
	books[2] = "Complete Golang"
	books[3] = "Complete Python"

	// print the type
	fmt.Printf("books  : %T\n", books)

	// print the elements
	fmt.Println("books  :", books)

	// print the elements in quoted string
	fmt.Printf("books  : %q\n", books)

	// print the elements along with their types
	fmt.Printf("books  : %#v\n", books)
}
