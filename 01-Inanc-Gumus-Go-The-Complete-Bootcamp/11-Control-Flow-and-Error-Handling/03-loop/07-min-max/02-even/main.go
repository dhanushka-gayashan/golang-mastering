package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	const usage = `Usage [min-number] [max-number]`

	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	min, minErr := strconv.Atoi(os.Args[1])
	max, maxErr := strconv.Atoi(os.Args[2])

	if minErr != nil || maxErr != nil {
		fmt.Println("Both values should be integer\n", usage)
		return
	}

	if min > max {
		fmt.Println("Min value should be less than Max\n", usage)
		return
	}

	var sum int
	result := ""

	for i := min; i <= max; i++ {

		if i%2 != 0 {
			continue
		}

		sum += i

		result = result + strconv.Itoa(i)

		if i < max {
			result = result + " + "
		}
	}

	fmt.Println(result+" = ", sum)
}
