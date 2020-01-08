package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	paths := os.Args[1:]
	if len(paths) == 0 {
		fmt.Println("Please provide directory paths")
		return
	}

	var directories []byte

main:
	for _, root := range paths {

		dir, err := ioutil.ReadDir(root)
		if err != nil {
			continue main
		}

		directories = append(directories, root...)
		directories = append(directories, '/', '\n')

		for _, file := range dir {
			if !file.IsDir() {
				directories = append(directories, '\t')
				directories = append(directories, file.Name()...)
				directories = append(directories, '/', '\n')
			}
		}

		directories = append(directories, '\n')
	}

	err := ioutil.WriteFile("ordered-list.txt", directories, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file created successfully")
}
