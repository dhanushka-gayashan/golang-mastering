package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	// Handle Errors
	switch {
	case len(args) < 5:
		fmt.Println("Please give me up to 5 numbers.")
		return
	case len(args) > 5:
		fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	// Filter numbers
	var sortedArray [5]int
	for i, rawNumber := range args {
		number, err := strconv.Atoi(rawNumber)

		if err != nil {
			continue
		}

		sortedArray[i] = number
	}

	// Sort the array -> https://youtu.be/nmhjrI-aW5o?t=7
	lastIndex := len(sortedArray) - 1

	//for i := 0; i < lastIndex; i++ {
	//	for j := 0; j < lastIndex - i; j++ {
	//		if sortedArray[j] > sortedArray[j + 1] {
	//			sortedArray[j], sortedArray[j + 1] = sortedArray[j + 1], sortedArray[j]
	//		}
	//	}
	//}

	for range sortedArray {
		for i, v := range sortedArray {
			if i < lastIndex && v > sortedArray[i+1] {
				sortedArray[i], sortedArray[i+1] = sortedArray[i+1], sortedArray[i]
			}
		}
	}

	fmt.Println(sortedArray)
}
