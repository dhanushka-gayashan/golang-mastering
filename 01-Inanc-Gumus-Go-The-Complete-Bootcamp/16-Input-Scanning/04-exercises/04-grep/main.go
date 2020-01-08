package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give a word....")
		return
	}

	query := args[0]

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		line := in.Text()
		present := strings.Contains(line, query)
		if present {
			fmt.Println(line)
		}
	}
}

// go run main.go dhanu < text.txt
