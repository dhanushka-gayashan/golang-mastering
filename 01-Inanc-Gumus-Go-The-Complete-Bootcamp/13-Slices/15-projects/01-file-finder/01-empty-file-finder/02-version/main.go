package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	var names []byte

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...) // ... convert string to byte slice
			names = append(names, '\n')
		}
	}

	err = ioutil.WriteFile("output.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", names)
}
