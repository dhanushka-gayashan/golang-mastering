package main

import (
	"fmt"
	"github.com/dhanukdg/Golang-Udemy-Todd-McLeod-Learn-How-To-Code-Google-Go/11-Testing-and-Benchmarking/03-example-test/calculation"
)

func main() {
	fmt.Println(calculation.Sum(2, 3, 4, 5))
	fmt.Println(calculation.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}
