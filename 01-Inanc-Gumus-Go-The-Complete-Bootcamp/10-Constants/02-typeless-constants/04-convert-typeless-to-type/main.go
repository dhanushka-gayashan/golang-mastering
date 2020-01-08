package main

import (
	"fmt"
	"time"
)

func main() {

	// Typeless Constant
	const c = 10

	// Go convert "c" into Duration type
	t := time.Second * c

	fmt.Println(t)
}
