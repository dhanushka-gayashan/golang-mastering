package main

import "fmt"

func main() {

	phone := map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}
	ph, ok := phone["dulin"]
	if ok {
		fmt.Println(ph)
	} else {
		fmt.Println("Value not found")
	}

	product := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}
	pr, ok := product[879401371]
	if ok {
		fmt.Println(pr)
	} else {
		fmt.Println("Value not found")
	}

	phones := map[string][]string{
		"bowen": []string{"202-555-0179"},
		"dulin": []string{"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": []string{"03489940240", "03489900120"},
	}
	phs, ok := phones["greco"]
	if ok {
		fmt.Println(phs)
	} else {
		fmt.Println("Value not found")
	}

	shopping := map[int]map[int]int{
		100: map[int]int{617841573: 4, 576872813: 2},
		101: map[int]int{576872813: 5, 657473833: 20},
		102: map[int]int{},
	}
	shp, ok := shopping[101]
	if ok {
		it, ok := shp[576872813]
		if ok {
			fmt.Println(it)
		} else {
			fmt.Println("Value not found")
		}
	} else {
		fmt.Println("Value not found")
	}
}
