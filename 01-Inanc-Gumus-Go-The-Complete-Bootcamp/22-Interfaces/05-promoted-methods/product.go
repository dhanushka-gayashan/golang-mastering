package main

import "fmt"

type product struct {
	title string
	price money
}

func (t *product) print() {
	fmt.Printf("%-15s: %s\n", t.title, t.price.string())
}

func (t *product) discount(ratio float64) {
	t.price *= money(1 - ratio)
}
