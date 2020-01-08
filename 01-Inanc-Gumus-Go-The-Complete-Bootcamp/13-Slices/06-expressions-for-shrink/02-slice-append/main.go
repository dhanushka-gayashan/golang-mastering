package main

import "fmt"

func main() {

	msg := []byte{'h', 'e', 'l', 'l', '0'}

	fmt.Println("msg", string(msg))

	_ = append(msg[:4], '!')
	fmt.Println("msg", string(msg))
}
