package main

import (
	"fmt"
	"strings"
)

func main() {

	// Trim - Both Side
	msg := `
	
	Hello Golang World
	
	`
	fmt.Println(msg)

	trimedMsg := strings.TrimSpace(msg)
	fmt.Println(trimedMsg)

	// Trim - Right and Left
	msg = "     Hello Golang World    "
	fmt.Println(msg)

	trimedMsg = strings.TrimRight(msg, " ")
	fmt.Println(trimedMsg)

	trimedMsg = strings.TrimLeft(trimedMsg, " ")
	fmt.Println(trimedMsg)
}
