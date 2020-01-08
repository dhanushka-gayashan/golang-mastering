package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arg := os.Args[1:]
	l := len(arg)
	if (l == 0) || l > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers)")
		return
	}

	var (
		sum   float64
		total float64
		nums  [5]float64
	)

	for i, n := range arg {
		j, err := strconv.ParseFloat(n, 64)

		if err != nil {
			continue
		}

		total++
		sum += j
		nums[i] = j
	}

	fmt.Println("Your numbers:", nums)
	fmt.Println("Average:", sum/total)
}
