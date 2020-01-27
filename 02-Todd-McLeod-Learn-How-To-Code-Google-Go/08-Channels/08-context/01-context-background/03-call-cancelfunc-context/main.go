package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t\t", ctx)
	fmt.Println("context error:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n\n", ctx)

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t\t", ctx)
	fmt.Println("context error:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n\n", ctx)

	cancel()

	fmt.Println("context:\t\t", ctx)
	fmt.Println("context error:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n\n", ctx)

	fmt.Println("cancel:\t\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n\n", cancel)
}
