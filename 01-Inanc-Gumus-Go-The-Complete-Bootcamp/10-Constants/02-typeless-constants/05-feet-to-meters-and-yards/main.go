package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	const (
		feetInMeter float64 = 0.3048
		feetInYards         = feetInMeter / 0.9144
	)

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * feetInMeter
	yards := math.Round(feet * feetInYards)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
