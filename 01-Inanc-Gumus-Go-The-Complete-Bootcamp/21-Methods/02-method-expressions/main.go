package main

func main() {

	var (
		mobydick = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris = game{title: "tetris", price: 5}
	)

	// Method Expression : Call a Methods through a Type
	book.print(mobydick)
	game.print(minecraft)
	game.print(tetris)
}

// go run .
