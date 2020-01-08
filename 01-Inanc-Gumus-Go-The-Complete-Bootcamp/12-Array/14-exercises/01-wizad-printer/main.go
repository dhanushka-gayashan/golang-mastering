package main

import (
	"fmt"
	"strings"
)

func main() {

	scientists := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton ", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	for i := range scientists {
		scientist := scientists[i]
		fmt.Printf("%-15s %-15s %-15s\n", scientist[0], scientist[1], scientist[2])

		if i == 0 {
			fmt.Println(strings.Repeat("=", 45))
		}
	}
}
