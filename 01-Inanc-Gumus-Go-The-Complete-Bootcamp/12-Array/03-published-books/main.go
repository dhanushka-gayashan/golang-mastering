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

	var published [len(books)]bool

	//Add published books
	published[0] = true
	published[len(books)-1] = true

	//Print Published Book's Title
	for i, ok := range published {
		if ok {
			fmt.Printf("+ %s\n", books[i])
		}
	}
}
