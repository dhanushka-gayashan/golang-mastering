package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fmt.Println("Scanned Text :", in.Text())
	}
}

// go run main.go < text.txt
