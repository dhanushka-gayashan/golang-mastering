package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	Report()

	millions := Read()

	millions = append([]int(nil), millions[len(millions)-10:]...)

	//last10 := make([]int, 10)
	//copy(last10, millions[len(millions)-10:])
	//millions = last10

	Report()

	fmt.Fprintln(ioutil.Discard, millions[0])
}
