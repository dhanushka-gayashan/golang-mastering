package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	houses := map[string][]string{
		"gryffindor": []string{"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  []string{"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  []string{"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  []string{"horace", "nigellus", "higgs", "scorpius"},
		"bobo":       []string{"wizardry", "unwanted"},
	}

	delete(houses, "bobo")

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	house, students := args[0], houses[args[0]]
	if students == nil {
		fmt.Printf("Sorry. I don't know anything about %q.\n", house)
		return
	}

	clone := append([]string(nil), students...)
	sort.Strings(clone)

	fmt.Printf("~~~ %s students ~~~\n\n", house)
	for _, student := range clone {
		fmt.Printf("\t+ %s\n", student)
	}
}
