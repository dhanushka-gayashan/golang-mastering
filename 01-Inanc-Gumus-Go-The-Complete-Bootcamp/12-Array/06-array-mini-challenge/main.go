package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("[your name]")
		return
	}

	moods := [...]string{
		"good",
		"happy",
		"awesome",
		"bad",
		"sad",
		"terrible",
	}

	name := args[0]

	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(len(moods))

	fmt.Printf("%s feels %s\n", name, moods[index])
}
