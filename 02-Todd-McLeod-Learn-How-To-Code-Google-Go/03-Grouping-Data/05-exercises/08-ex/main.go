package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Being evil", "Ice cream", "Sunsets"},
	}

	for k, v := range m {
		fmt.Println(k)
		for i := 0; i < len(v); i++ {
			fmt.Println("\t", v[i])
		}
		fmt.Println()
	}
}
