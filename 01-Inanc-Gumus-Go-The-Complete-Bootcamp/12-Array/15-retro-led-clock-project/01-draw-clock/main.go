package main

import (
	"fmt"
	"time"
)

func main() {

	// Step 1 - Define Placeholder Underlying Type
	type placeholder [5]string

	// Step 2 - Define Digits
	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	// Step 3 - Define Digits Array
	digits := [10][5]string{zero, one, two, three, four, five, six, seven, eight, nine}

	// Step 4 - Print Side by Side
	// for i:=0; i <5; i++
	//for line:= range digits[0]{
	//	for digit := range digits {
	//		fmt.Print(digits[digit][line], "  ")
	//	}
	//	fmt.Println()
	//}

	// Step 5 - Get the current time
	now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	// Step 6 - Define Clock Array
	clock := [...]placeholder{
		digits[hour/10], digits[hour%10],
		colon,
		digits[min/10], digits[min%10],
		colon,
		digits[sec/10], digits[sec%10],
	}

	// Step 7 - Print Clock
	for line := range clock[0] {
		for _, digit := range clock {
			fmt.Print(digit[line], "  ")
		}
		fmt.Println()
	}

}
