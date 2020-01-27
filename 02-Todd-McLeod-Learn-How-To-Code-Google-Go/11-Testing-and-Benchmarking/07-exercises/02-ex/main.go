package main

import (
	"fmt"
	"github.com/dhanukdg/Golang-Udemy-Todd-McLeod-Learn-How-To-Code-Google-Go/11-Testing-and-Benchmarking/07-exercises/02-ex/quote"
	"github.com/dhanukdg/Golang-Udemy-Todd-McLeod-Learn-How-To-Code-Google-Go/11-Testing-and-Benchmarking/07-exercises/02-ex/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
