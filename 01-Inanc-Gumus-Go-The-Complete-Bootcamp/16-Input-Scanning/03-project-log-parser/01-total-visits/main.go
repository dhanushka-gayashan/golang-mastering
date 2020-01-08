package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	sum := map[string]int{}

	for in.Scan() {

		fields := strings.Fields(in.Text())

		domain := fields[0]

		visit, _ := strconv.Atoi(fields[1])

		sum[domain] += visit
	}

	fmt.Printf("%-30s %-10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for domain, visits := range sum {
		fmt.Printf("%-30s %-10s\n", domain, visits)
	}
}
