package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Fatal functions call os.Exit(1) after writing the log message ....
	// defer() won't run
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln("log.Fatalln() -> Error happen:", err)
	}

 	defer foo()
}

func foo() {
	fmt.Println("Going to call in defer()")
}
