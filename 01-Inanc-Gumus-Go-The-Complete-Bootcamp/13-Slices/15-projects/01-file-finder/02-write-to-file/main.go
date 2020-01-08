package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(args)

	var items []byte
	for _, item := range args {
		item = item + "\n"
		items = append(items, item...)
	}

	err := ioutil.WriteFile("ordered-list.txt", items, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File has been created successfully")
}
