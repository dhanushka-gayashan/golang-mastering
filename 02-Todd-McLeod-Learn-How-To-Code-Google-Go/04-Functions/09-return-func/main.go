package main

import "fmt"

func main() {

	f1 := foo()
	f1()


	f2 := bar()
	fmt.Println(f2())
	fmt.Println(bar()())


	f3 := zoo()
	fmt.Println(f3("Nick"))
	fmt.Println(zoo()("Nick"))
}

func foo() func(){

	return func() {
		fmt.Println("Got from foo")
	}
}

func bar() func() string{

	return func() string {
		return "Got from bar"
	}
}

func zoo() func(name string) string {

	return func(name string) string {
		return "Hi, " + name
	}
}
