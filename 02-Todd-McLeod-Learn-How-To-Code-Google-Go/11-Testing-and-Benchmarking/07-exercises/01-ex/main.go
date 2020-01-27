package main

import (
	"fmt"
	"github.com/dhanukdg/Golang-Udemy-Todd-McLeod-Learn-How-To-Code-Google-Go/11-Testing-and-Benchmarking/07-exercises/01-ex/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
