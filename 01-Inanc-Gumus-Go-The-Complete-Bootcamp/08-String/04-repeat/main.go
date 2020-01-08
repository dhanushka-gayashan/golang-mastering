package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	msg := strings.ToUpper(os.Args[1])
	marks := strings.Repeat("!", len(msg))
	msg = marks + msg + marks
	fmt.Println(msg)
}
