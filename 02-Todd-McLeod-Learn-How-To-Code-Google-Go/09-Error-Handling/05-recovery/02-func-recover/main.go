package main

import "fmt"

func main() {
	foo()
	fmt.Println("Returned normally from Foo()")
}

func foo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in foo()", r)
		}
	}()

	fmt.Println("Calling Bar()")

	bar(0)

	// This won't execute, because bar() got panic and it's need to be recovered
	fmt.Println("Returned normally from Bar()")
}

func bar(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("\nBar() Panic -> value: %v", i))
	}

	defer fmt.Println("Bar() Defer -> value:", i)

	fmt.Println("Bar() Println -> value:", i)

	bar(i + 1)
}
