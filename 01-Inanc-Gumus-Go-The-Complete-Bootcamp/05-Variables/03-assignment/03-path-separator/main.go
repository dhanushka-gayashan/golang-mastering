package main

import (
	"fmt"
	"path"
)

func main() {

	dir, file := path.Split("css/main.css")

	fmt.Println("Dir is ", dir)
	fmt.Println("File is ", file)
}
