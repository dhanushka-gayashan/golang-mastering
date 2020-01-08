package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	in.Scan()

	fmt.Println("Scanned Text : ", in.Text())
	fmt.Println("Scanned Bytes : ", in.Bytes())
}
