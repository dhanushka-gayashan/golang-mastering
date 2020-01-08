package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give a path....")
		return
	}

	custPath := strings.ToLower(os.Args[1])

	//pathEnv, _ := os.LookupEnv("path")
	//pathArray := strings.Split(pathEnv, ";")
	pathArray := filepath.SplitList(os.Getenv("path"))

	for i, p := range pathArray {

		if strings.Contains(strings.ToLower(p), custPath) {
			fmt.Printf("#-2%d : %q\n", i, p)
		}
	}
}
