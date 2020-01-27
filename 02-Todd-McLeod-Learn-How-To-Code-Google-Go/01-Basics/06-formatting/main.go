package main

import "fmt"

var x = 50

func main() {

	fmt.Printf("Real Value : %v\n", x)
	fmt.Printf("Type of Value : %T\n", x)
	fmt.Printf("Binary of Value : %b\n", x)
	fmt.Printf("Hexadecimal of Value : %x\n", x)
	fmt.Printf("Hexadecimal of Value : %#x\n", x)

	y := 40
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	str := fmt.Sprintf("%#x\t%b\t%x", y, y, y)
	fmt.Println(str)
}
