package main

import "fmt"

func main() {

	n, m := 10, 20

	fmt.Printf("Before call multiply() n : %d  m : %d\n", n, m)

	result := multiply(n, m)
	fmt.Println("Result of multiply() :", result)

	fmt.Printf("After call multiply() n : %d  m : %d\n", n, m)
}

func multiply(a, b int) int {

	fmt.Printf("Inside multiply() - Before multiply a : %d  b : %d\n", a, b)

	// Go declare Local Variables for Parameters
	a *= b

	fmt.Printf("Inside multiply() - After multiply a : %d  b : %d\n", a, b)

	return a
}
