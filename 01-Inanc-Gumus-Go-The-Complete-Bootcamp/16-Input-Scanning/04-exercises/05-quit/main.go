package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	words := map[string]bool{}

	for in.Scan() {
		word := strings.ToLower(in.Text())
		if words[word] {
			fmt.Println("TWICE!")
			return
		}
		words[word] = true
	}
}
