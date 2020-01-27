package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Hello Golang")

	_, _ = fmt.Fprintln(os.Stdout, "Hello Golang")

	_, _ = io.WriteString(os.Stdout, "Hello Golang")
}
