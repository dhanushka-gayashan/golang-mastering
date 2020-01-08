package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me a moth name")
		return
	}

	month := strings.ToLower(args[1])
	days := 28

	if month == "january" || month == "march" || month == "may" || month == "july" || month == "august" || month == "october" || month == "december" {
		days = 31
	} else if month == "april" || month == "june" || month == "september" || month == "november" {
		days = 30
	} else if month == "february" {
		year := time.Now().Year()
		if (year%4 == 0) && (year%100 != 0 && year%400 == 0) {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", month)
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
