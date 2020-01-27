package main

import (
	"context"
	"fmt"
)

type key string

func main() {
	ctx := context.WithValue(context.Background(), key("language"), "Go")
	foo(ctx, "language")

	ctx = context.WithValue(ctx, key("dog"), "Gaston")
	foo(ctx, "dog")

	foo(ctx, "color")
}

func foo(ctx context.Context, s string) {
	if v := ctx.Value(s); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found:", s)
}
