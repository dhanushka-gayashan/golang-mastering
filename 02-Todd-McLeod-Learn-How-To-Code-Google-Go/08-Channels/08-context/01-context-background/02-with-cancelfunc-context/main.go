package main

import (
	"context"
	"fmt"
)

func main() {
	ctx1 := context.Background()

	fmt.Println("context:\t\t", ctx1)
	fmt.Println("context error:\t", ctx1.Err())
	fmt.Printf("context type:\t%T\n\n", ctx1)

	ctx2, _ := context.WithCancel(ctx1)

	fmt.Println("context:\t\t", ctx2)
	fmt.Println("context error:\t", ctx2.Err())
	fmt.Printf("context type:\t%T\n", ctx2)
}
