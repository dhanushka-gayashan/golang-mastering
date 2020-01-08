package main

import (
	"github.com/dhanukdg/Golang-Udemy-Inanc-Gumus-Go-The-Complete-Bootcamp/23-Famous-Interfaces/03-sorter-interface/01-default-implementation"
	"sort"
)

func main() {

	l := _1_default_implementation.list{
		{title: "moby dick", price: 10, released: _1_default_implementation.toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: _1_default_implementation.toTimestamp("733622400")},
		{title: "hobbit", price: 25},
	}

	sort.Sort(l)

	sort.Sort(sort.Reverse(l))
}
