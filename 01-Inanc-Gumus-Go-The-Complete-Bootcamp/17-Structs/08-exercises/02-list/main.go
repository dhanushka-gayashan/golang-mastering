package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {

	games := []game{
		{item: item{id: 1, name: "god of war", price: 50}, genre: "action adventure"},
		{item: item{id: 2, name: "x-com 2", price: 40}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	in := bufio.NewScanner(os.Stdin)

	fmt.Println("Available options: List l | Item <1/2/3> | Quit q")

	for in.Scan() {

		switch in.Text() {
		case "l":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)
			}
		case "q":
			os.Exit(0)
		case "1", "2", "3":
			n, _ := strconv.Atoi(in.Text())
			g := games[n-1]
			fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)
		default:
			fmt.Println("Invalid Option...")
		}
	}
}
