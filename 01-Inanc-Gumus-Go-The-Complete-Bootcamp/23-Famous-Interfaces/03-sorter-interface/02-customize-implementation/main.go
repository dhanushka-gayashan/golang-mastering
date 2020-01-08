package main

import (
	"sort"
)

func main() {

	l := list{
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
	}

	sort.Sort(byReleaseDate(l))

	sort.Sort(sort.Reverse(byReleaseDate(l)))
}
