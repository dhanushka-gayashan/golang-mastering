package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	data := "2 4 6 1 3 5"

	strData := strings.Split(data, " ")

	var intData []int
	for i := range strData {
		n, _ := strconv.Atoi(strData[i])
		intData = append(intData, n)
	}
	fmt.Println("intData[]", intData)

	var even []int
	for i := range intData {
		n := intData[i]
		if n%2 == 0 {
			even = append(even, n)
		}
	}
	fmt.Println("even[]", even)

	var odd []int
	for i := range intData {
		n := intData[i]
		if n%2 != 0 {
			odd = append(odd, n)
		}
	}
	fmt.Println("odd[]", odd)

	middle := intData[2:4]
	fmt.Println("middle[]", middle)

	first2 := intData[:2]
	fmt.Println("first2[]", first2)

	last2 := intData[len(intData)-2:]
	fmt.Println("last2[]", last2)

	evenLast1 := even[len(even)-1:]
	fmt.Println("evenLast1[]", evenLast1)

	oddLast2 := odd[len(odd)-2:]
	fmt.Println("oddLast2[]", oddLast2)
}
