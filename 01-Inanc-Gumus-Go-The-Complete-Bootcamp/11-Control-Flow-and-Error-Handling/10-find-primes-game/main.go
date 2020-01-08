package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

main:
	for _, q := range os.Args[1:] {

		n, err := strconv.Atoi(q)

		if err != nil {
			continue
		}

		// https://stackoverflow.com/questions/1801391/how-to-create-the-most-compact-mapping-n-%e2%86%92-isprimen-up-to-a-limit-n/1801446#1801446
		switch {
		case n == 2 || n == 3: //Prime
			fmt.Printf("%d ", n)
		case n <= 1 || n%2 == 0 || n%3 == 0: //Not a Prime
			continue
		}

		for i, w := 5, 2; i*i <= n; {
			// not a prime
			if n%i == 0 {
				continue main
			}

			i += w
			w = 6 - w
		}

		// all checks ok: it's a prime
		fmt.Print(n, " ")
	}
	fmt.Println()
}
