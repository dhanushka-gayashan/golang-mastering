package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(args)

	var items []byte
	for i, item := range args {
		items = strconv.AppendInt(items, int64(i+1), 10)
		items = append(items, '.', ' ')
		items = append(items, item...)
		items = append(items, '\n')
	}

	err := ioutil.WriteFile("ordered-list.txt", items, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File has been created successfully")
}
