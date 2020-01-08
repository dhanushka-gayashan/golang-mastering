package main

import "fmt"

func main() {

	sliceable := []byte{'f', 'u', 'l', 'l'}

	show(sliceable[0:3])

	show(sliceable[0:3:3])

	show(sliceable[0:2:2])

	show(sliceable[0:1:1])

	show(sliceable[1:3:3])

	show(sliceable[2:3:3])

	show(sliceable[2:3:4])

	show(sliceable[4:4:4]) // When append new element to this slice, create new underlying array, because Capacity = 0
}

func show(s []byte) {
	fmt.Printf("Length : %d Capacity : %d Slice %s\n", len(s), cap(s), s)
}
