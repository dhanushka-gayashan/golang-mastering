package main

type item interface {
	print()
	discount(ratio float64)
}
