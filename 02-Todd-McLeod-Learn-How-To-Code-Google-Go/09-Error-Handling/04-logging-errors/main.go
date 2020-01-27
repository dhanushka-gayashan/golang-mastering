package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//log file
	f, err := os.Create("app.log")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
	defer f2.Close()

	fmt.Println("Check the app.log file in the directory...")
}
