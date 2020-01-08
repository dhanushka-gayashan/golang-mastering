package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	max, _ := strconv.Atoi(os.Args[1])

	var uniques []int

main:
	for len(uniques) < max {
		n := rand.Intn(max + 1)

		for _, u := range uniques {
			if n == u {
				continue main
			}
		}

		uniques = append(uniques, n)
	}

	fmt.Println("uniques : ", uniques)

	sort.Ints(uniques)
	fmt.Println("sorted uniques : ", uniques)

	// Sort -> Need Slice -> Array to Slice -> [5]number -> number[:]
	nums := [5]int{5, 4, 3, 2, 1}
	sort.Ints(nums[:])
	fmt.Println("\nnums:", nums)
}
