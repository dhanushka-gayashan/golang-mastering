package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		headers                    []string
		houses                     []string
		locations                  []string
		sizes, beds, baths, prices []int
		averages                   []float64

		sizesSum  int
		bedsSum   int
		bathsSum  int
		pricesSum int
	)

	headers = strings.Split(header, separator)

	houses = strings.Split(data, "\n")

	for _, h := range houses {
		house := strings.Split(h, separator)
		locations = append(locations, house[0])

		size, _ := strconv.Atoi(house[1])
		sizes = append(sizes, size)
		sizesSum += size

		bed, _ := strconv.Atoi(house[2])
		beds = append(beds, bed)
		bedsSum += bed

		bath, _ := strconv.Atoi(house[3])
		baths = append(baths, bath)
		bathsSum += bath

		price, _ := strconv.Atoi(house[4])
		prices = append(prices, price)
		pricesSum += price
	}

	divider := float64(len(houses))
	averages = append(averages, float64(sizesSum)/divider)
	averages = append(averages, float64(bedsSum)/divider)
	averages = append(averages, float64(bathsSum)/divider)
	averages = append(averages, float64(pricesSum)/divider)

	for i := range headers {
		fmt.Printf("%-15s", headers[i])
	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 70))

	for i := range locations {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}

	fmt.Printf("%s\n", strings.Repeat("=", 70))

	fmt.Printf("%-15s%-15.2f%-15.2f%-15.2f%-15.2f\n", "", averages[0], averages[1], averages[2], averages[3])
}
