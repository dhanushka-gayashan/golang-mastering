package main

import "fmt"

func main() {

	msg := []byte{'h', 'e', 'l', 'l', '0'}

	fmt.Println("msg", string(msg))

	fmt.Println("msg[0:1]", string(msg[0:1]))
	fmt.Println("msg[0:2]", string(msg[0:2]))
	fmt.Println("msg[0:3]", string(msg[0:3]))
	fmt.Println("msg[0:4]", string(msg[0:4]))
	fmt.Println("msg[0:5]", string(msg[0:5]))

	fmt.Println("msg[0:]", string(msg[0:]))
	fmt.Println("msg[:5]", string(msg[:5]))
	fmt.Println("msg[:]", string(msg[:]))

	fmt.Println("msg[1:4]", string(msg[1:4]))
	fmt.Println("msg[1:5]", string(msg[1:5]))
	fmt.Println("msg[1:]", string(msg[1:]))
}
