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

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	wBooks[0] = books[0]

	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	fmt.Printf("\nWinter Books : %#v\n", wBooks)
	fmt.Printf("\nSummer Books : %#v\n", sBooks)
}
