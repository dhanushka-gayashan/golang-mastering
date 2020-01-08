package main

import (
	"fmt"
	"sort"
)

func main() {

	grades := [...]float64{40, 20, 30, 50, 60, 70}

	front := grades[0:3]

	sort.Float64s(front)

	fmt.Println("grades:", grades)
}
