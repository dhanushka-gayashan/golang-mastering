package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Being evil", "Ice cream", "Sunsets"},
		"batman": []string{"Catwoman", "Jocker", "Robbin"},
	}
	fmt.Println(m)

	fmt.Println()

	if _, ok := m["batman"]; ok {
		delete(m, "batman")
	}
	fmt.Println(m)
}
