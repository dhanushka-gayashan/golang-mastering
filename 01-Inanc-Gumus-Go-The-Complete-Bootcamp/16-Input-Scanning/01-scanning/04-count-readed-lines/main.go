package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//Check Error work
	//os.Stdin.Close()

	in := bufio.NewScanner(os.Stdin)

	var lines int

	for in.Scan() {
		lines++
	}
	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}
}

// go run main.go < text.txt
