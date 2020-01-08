package main

import "fmt"

func main() {

	store := list{
		&book{product{"moby dick", 10}, 118281600},
		&book{product{"odyssey", 15}, "733622400"},
		&book{product{"hobbit", 25}, nil},
		&puzzle{product{"rubik's cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&toy{product{"yoda", 150}},
	}
	store.discount(.5)
	store.print()


	//Embedded Values and Promoted Methods
	t := &toy{product{"yoda", 150}}
	fmt.Printf("%#v\n", t)
	t.print()


	//Embedded Values and Promoted Methods
	b := &book{product{"moby dick", 10}, 118281600}
	fmt.Printf("%#v\n", b)
	b.print()
	b.product.print()
}
