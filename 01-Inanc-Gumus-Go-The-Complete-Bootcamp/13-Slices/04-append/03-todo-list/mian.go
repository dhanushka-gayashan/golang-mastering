package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {

	var todo []string

	todo = append(todo, "sing")
	todo = append(todo, "run")
	todo = append(todo, "code", "play", "sleep")

	tomorrow := []string{"movie", "gaming"}
	todo = append(todo, tomorrow...)

	s.Show("todo", todo)
}
