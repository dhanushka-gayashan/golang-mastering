package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	//Only Buffer last scan line
	in.Scan()
	in.Scan()
	in.Scan()

	fmt.Println("Scanned Text : ", in.Text())
	fmt.Println("Scanned Text : ", in.Text())
	fmt.Println("Scanned Text : ", in.Text())
}
