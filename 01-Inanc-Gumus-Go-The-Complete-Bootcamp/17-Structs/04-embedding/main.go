package main

import "fmt"

func main() {

	// Embedding == Composition (Book contain Text)
	type text struct {
		title string
		words int
	}

	type book struct {
		text text
		isbn string
	}

	moby := book{
		text: text{"moby dick", 206053},
		isbn: "102030",
	}
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.text.title, moby.text.words, moby.isbn)

	// Anonymous fields
	type article struct {
		text
		link string
	}

	greatGo := article{
		text: text{"Great Golang", 400000},
		link: "http://golang.org/books/great.pdf",
	}
	fmt.Printf("%s has %d words (link: %s)\n", greatGo.text.title, greatGo.text.words, greatGo.link)

	// Change Field Values
	greatGo.text.words = 500000
	fmt.Printf("%s has %d words (link: %s)\n", greatGo.text.title, greatGo.text.words, greatGo.link)

	// Type Less
	greatGo.words++
	fmt.Printf("%s has %d words (link: %s)\n", greatGo.title, greatGo.words, greatGo.link)
}
