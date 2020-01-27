package main

import "fmt"

func main() {

	s:= "Hello, Welcome to Golang world"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	//UTF-8 Represent of String
	//UTF-8 uses Hexadecimal
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
}
