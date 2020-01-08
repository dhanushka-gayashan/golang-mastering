package main

import "fmt"

func main() {

	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte

	for _, word := range words {
		b := []byte(word)
		fmt.Println(b)
		bwords = append(bwords, b)
	}

	for _, word := range bwords {
		s := string(word)
		fmt.Println(s)
	}
}
