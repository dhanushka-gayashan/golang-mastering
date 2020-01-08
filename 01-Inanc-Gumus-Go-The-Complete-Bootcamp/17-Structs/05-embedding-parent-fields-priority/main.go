package main

import "fmt"

func main() {

	type text struct {
		title string
		words int
	}

	type book struct {
		text  text
		isbn  string
		title string
	}

	book1 := book{
		text: text{"Go World I", 1000},
		isbn: "11223345",
	}
	fmt.Printf("Book 1 Title : %s\n", book1.title)
	fmt.Printf("Book 1 Text -> Title : %s\n", book1.text.title)

	book2 := book{
		text:  text{"Go World II", 2000},
		isbn:  "66778899",
		title: "Grate Way To Start",
	}
	fmt.Printf("Book 2 Title : %s\n", book2.title)
}
