package main

import "fmt"

func main() {

	fmt.Println("doubles:")

	words := []string{1022: ""}
	fmt.Println("l:", len(words), "c:", cap(words))

	words = append(words, "boom!")
	fmt.Println("l:", len(words), "c:", cap(words))

	// ------------------------------------------------

	// grows 1.25 (cap is 1024, hence >= 1024)
	fmt.Println("\ngrows %25:")

	words = []string{1023: ""}
	fmt.Println("l:", len(words), "c:", cap(words))

	words = append(words, "boom!")
	fmt.Println("l:", len(words), "c:", cap(words))
}
